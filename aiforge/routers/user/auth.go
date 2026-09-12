// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2018 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"code.gitea.io/gitea/modules/auth/wechat"

	"code.gitea.io/gitea/modules/notification"

	"code.gitea.io/gitea/modules/redis/redis_client"
	"code.gitea.io/gitea/modules/redis/redis_key"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/login_service"
	"code.gitea.io/gitea/services/sign_up_service"

	"github.com/gomodule/redigo/redis"

	"code.gitea.io/gitea/modules/slideimage"

	phoneService "code.gitea.io/gitea/services/phone"

	"code.gitea.io/gitea/modules/labelmsg"

	"code.gitea.io/gitea/modules/phone"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth"
	"code.gitea.io/gitea/modules/auth/oauth2"
	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/eventsource"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/password"
	"code.gitea.io/gitea/modules/recaptcha"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"
	"code.gitea.io/gitea/modules/util"
	"code.gitea.io/gitea/services/externalaccount"
	"code.gitea.io/gitea/services/mailer"

	"gitea.com/macaron/captcha"
	"github.com/markbates/goth"
	"github.com/tstranex/u2f"
)

const (
	// tplMustChangePassword template for updating a user's password
	tplMustChangePassword = "user/auth/change_passwd"
	// tplSignIn template for sign in page
	tplSignIn base.TplName = "user/auth/signin"
	// tplSignIn template for sign in page
	tplSignInCloudBrain base.TplName = "user/auth/signin_cloud_brain"
	tplSignInPhone      base.TplName = "user/auth/signin_phone"
	tplSignInWeChat     base.TplName = "user/auth/signin_wechat"
	tplSignUpScanCode   base.TplName = "user/auth/signup_scancode"
	// tplSignUp template path for sign up page
	tplSignUp base.TplName = "user/auth/signup"
	// TplActivate template path for activate user
	TplActivate            base.TplName = "user/auth/activate"
	tplForgotPassword      base.TplName = "user/auth/forgot_passwd"
	tplForgotPasswordPhone base.TplName = "user/auth/forgot_passwd_phone"
	tplResetPassword       base.TplName = "user/auth/reset_passwd"
	tplTwofa               base.TplName = "user/auth/twofa"
	tplTwofaScratch        base.TplName = "user/auth/twofa_scratch"
	tplLinkAccount         base.TplName = "user/auth/link_account"
	tplLinkAccountNew      base.TplName = "user/auth/link_account_new"
	tplU2F                 base.TplName = "user/auth/u2f"
	tplGuestVerify         base.TplName = "user/auth/guest_verify"
)

// AutoSignIn reads cookie and try to auto-login.
func AutoSignIn(ctx *context.Context) (bool, error) {
	if !models.HasEngine {
		return false, nil
	}
	uname := ctx.GetCookie(setting.CookieUserName)
	if len(uname) == 0 {
		return false, nil
	}
	isSucceed := false
	defer func() {
		if !isSucceed {
			log.Trace("auto-login cookie cleared: %s", uname)
			ctx.SetCookie(setting.CookieUserName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
			ctx.SetCookie(setting.CookieRememberName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
		}
	}()

	u, err := models.GetUserByName(uname)
	if err != nil {
		if !models.IsErrUserNotExist(err) {
			return false, fmt.Errorf("GetUserByName: %v", err)
		}
		return false, nil
	}
	if val, ok := ctx.GetSuperSecureCookie(
		base.EncodeMD5(u.Rands+u.Passwd), setting.CookieRememberName); !ok || val != u.Name {
		return false, nil
	}
	isSucceed = true

	// Set session IDs
	if err := ctx.Session.Set("uid", u.ID); err != nil {
		return false, err
	}
	if err := ctx.Session.Set("uname", u.Name); err != nil {
		return false, err
	}
	if err := ctx.Session.Release(); err != nil {
		return false, err
	}
	log.Info("Auto login succeed.")
	models.SaveLoginInfoToDb(ctx.Req.Request, u)
	ctx.SetCookie(setting.CSRFCookieName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	return true, nil
}

func checkAutoLogin(ctx *context.Context) bool {
	// Check auto-login.
	isSucceed, err := AutoSignIn(ctx)
	if err != nil {
		ctx.ServerError("AutoSignIn", err)
		return true
	}
	redirectTo := ctx.Query("redirect_to")
	if len(redirectTo) > 0 {
		ctx.SetCookie("redirect_to", redirectTo, 0, setting.AppSubURL, "", setting.SessionConfig.Secure, true)
	} else {
		redirectTo = ctx.GetCookie("redirect_to")
	}

	if isSucceed {
		isCourse := ctx.QueryBool("course")
		ctx.SetCookie("redirect_to", "", -1, setting.AppSubURL, "", setting.SessionConfig.Secure, true)
		if redirectTo == "" && isCourse {
			redirectToCourse := setting.AppSubURL + "/" + setting.Course.OrgName
			ctx.RedirectToFirst(redirectToCourse)
		} else {
			ctx.RedirectToFirst(redirectTo, setting.AppSubURL+string(setting.LandingPageURL))
		}
		return true
	}

	return false
}

func getActivityTpl() string {
	// Avoid blocking login page render on external promote repo availability.
	return ""
}

// SignIn render sign in page
func SignIn(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_in")
	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}

	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["Title"] = ctx.Tr("sign_in")
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/login"
	ctx.Data["PageIsSignIn"] = true
	ctx.Data["IsCourse"] = ctx.QueryBool("course")
	ctx.Data["PageIsLogin"] = true
	ctx.Data["EnableSSPI"] = models.IsSSPIEnabled()
	ctx.Data["EnableCloudBrain"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	setRSAContext(ctx)

	ctx.Data["ActivityTpl"] = getActivityTpl()
	ctx.HTML(200, tplSignIn)
}

func GetSignUpWechatQRCode(ctx *context.Context) {
	sharedUser := ctx.Query("sharedUser")
	dataMap := map[string]interface{}{}
	if sharedUser != "" {
		dataMap["sharedUser"] = sharedUser
	}
	res, err := sign_up_service.GetSignUpWechatQRCode(dataMap)
	if err != nil {
		log.Error("GetSignUpWechatQRCode err.%+v", err)
		ctx.JSON(http.StatusOK, response.NewBizError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(res))

}

func GetLoginWechatQRCode(ctx *context.Context) {
	res, err := login_service.GetLoginWechatQRCode()
	if err != nil {
		log.Error("GetSignUpWechatQRCode err.%+v", err)
		ctx.JSON(http.StatusOK, response.NewBizError(err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessWithData(res))

}

func LoginByWechat(ctx *context.Context, form auth.WechatForm) {
	loginId := form.LoginId
	loginVal, _ := redis_client.Get(redis_key.LoginWechatKey(loginId))
	if loginVal == "" {
		log.Error("LoginWechatKey not exits.loginId=%s", loginId)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.QR_CODE_EXPIRED, ctx))
		return
	}
	loginCache := new(models.LoginWechatCache)
	json.Unmarshal([]byte(loginVal), loginCache)
	userId := loginCache.UserId
	if userId <= 0 {
		log.Error("userId in login wechat cache is not correct.loginId=%s loginCache=%s", loginId, loginVal)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.WECHAT_NOT_BOUND, ctx))
		return
	}
	u, err := models.GetUserByID(userId)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			log.Error("userId not exists.loginId=%s loginCache=%s", loginId, loginVal)
			ctx.JSON(http.StatusOK, response.OuterTrBizError(response.WECHAT_NOT_BOUND, ctx))
			return
		}
		log.Error("GetUserByID err .loginId=%s loginCache=%s err=%v", loginId, loginVal, err)
		ctx.JSON(http.StatusOK, response.OuterTrBizError(response.NewBizError(err), ctx))
		return
	}

	models.SaveLoginInfoToDb(ctx.Req.Request, u)

	handleSignIn(ctx, u, false)
	redis_client.Del(redis_key.LoginWechatKey(loginId))
}

func setRSAContext(ctx *context.Context) {
	ctx.Data["LoginEncrypt"] = setting.LoginEncrypt
	if setting.LoginEncrypt {
		privateKey, err := readKeyFile(setting.CustomPath + "/conf/publickey.pem")
		if err == nil {
			ctx.Data["RSAPublicKey"] = string(privateKey)
		}
	} else {
		ctx.Data["RSAPublicKey"] = ""
	}
}

func GetUserOtherInfo(ctx *context.Context) {
	userName := ctx.Query("userName")
	phone := ctx.Query("phone")
	var err error
	var user *models.User
	if userName != "" {
		user, err = models.GetUserByName(userName)
	}
	if phone != "" {
		user, err = models.GetUserByPhoneNumber(phone)
	}
	re := make(map[string]interface{}, 0)
	if err != nil || user == nil {
		re["result_code"] = "-1"
		ctx.JSON(200, re)
	} else {
		re["result_code"] = "0"
		re["agree"] = models.GetNewAgreeByUID(user.ID)
		ctx.JSON(200, re)
	}
}

func SaveUserOtherInfo(ctx *context.Context) {
	userName := ctx.Query("userName")
	phone := ctx.Query("phone")
	var err error
	var user *models.User
	if userName != "" {
		user, err = models.GetUserByName(userName)
	}
	if phone != "" {
		user, err = models.GetUserByPhoneNumber(phone)
	}
	re := make(map[string]interface{}, 0)
	if err != nil || user == nil {
		re["result_code"] = "-1"
		ctx.JSON(200, re)
	} else {
		re["result_code"] = "0"
		models.SaveUserOtherInfoToDb(user.ID)
		ctx.Data["IsNewAgree"] = true
		ctx.JSON(200, re)
	}
}

// SignInCloudBrain render sign in page
func SignInCloudBrain(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_in")

	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}
	setRSAContext(ctx)

	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/login/cloud_brain"
	ctx.Data["PageIsSignIn"] = true
	ctx.Data["PageIsCloudBrainLogin"] = true
	ctx.Data["EnableCloudBrain"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()

	ctx.HTML(200, tplSignInCloudBrain)
}

func SignInWeChat(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_in")
	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}
	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["PageIsWeChatLogin"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()

	ctx.HTML(200, tplSignInWeChat)
}

func GuestVerify(ctx *context.Context) {
	ctx.HTML(200, tplGuestVerify)
}

func SignInPhone(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_in")
	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}
	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["PageIsPhoneLogin"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()

	ctx.HTML(200, tplSignInPhone)
}

func SignInPhonePost(ctx *context.Context, form auth.PhoneNumberCodeForm) {
	ctx.Data["Title"] = ctx.Tr("sign_in")
	ctx.Data["PageIsPhoneLogin"] = true
	ctx.Data["IsCourse"] = ctx.QueryBool("course")
	ctx.Data["EnableCloudBrain"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()

	if ctx.HasError() {
		ctx.HTML(200, tplSignInPhone)
		return
	}

	if !phoneService.IsVerifyCodeRight(strings.TrimSpace(form.PhoneNumber), strings.TrimSpace(form.VerifyCode)) {
		ctx.RenderWithErr(ctx.Tr("phone.verify_code_fail"), tplSignInPhone, &form)
		return
	}

	u, err := models.GetUserByPhoneNumber(strings.TrimSpace(form.PhoneNumber))

	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.RenderWithErr(ctx.Tr("form.username_password_incorrect"), tplSignInPhone, &form)
			log.Info("Failed authentication attempt for %s from %s", form.PhoneNumber, ctx.RemoteAddr())
		} else {
			ctx.ServerError("UserSignIn", err)
		}
		return
	}

	models.SaveLoginInfoToDb(ctx.Req.Request, u)

	handleSignIn(ctx, u, form.Remember)
}

func SignInPostAPI(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_in")
	UserName := ctx.Query("UserName")
	Password := ctx.Query("Password")
	log.Info("u=" + UserName)
	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["Title"] = ctx.Tr("sign_in")
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/login"
	ctx.Data["PageIsSignIn"] = true
	ctx.Data["PageIsLogin"] = true
	ctx.Data["IsCourse"] = ctx.QueryBool("course")
	ctx.Data["EnableSSPI"] = models.IsSSPIEnabled()

	if ctx.HasError() {
		ctx.HTML(200, tplSignIn)
		return
	}
	u, err := models.UserSignIn(UserName, Password)
	if err != nil {
		log.Info("login failed.UserName=" + UserName + " Password=" + Password)
		ctx.ServerError("UserSignIn", err)
		return
	}
	models.SaveLoginInfoToDb(ctx.Req.Request, u)
	// If this user is enrolled in 2FA, we can't sign the user in just yet.
	// Instead, redirect them to the 2FA authentication page.
	//handleSignInFull(ctx, u, form.Remember, false)
	handleSignInFullNotRedirect(ctx, u, true, false)
}

func readKeyFile(keypath string) ([]byte, error) {
	privatekeyfile, err := os.Open(keypath)
	if err != nil {
		return nil, err
	}
	defer privatekeyfile.Close()
	//get privatekeyfile content
	privatekeyinfo, _ := privatekeyfile.Stat()
	buf := make([]byte, privatekeyinfo.Size())
	privatekeyfile.Read(buf)
	return buf, nil
}

func RsaDecrypt(password string, privatekeypath string) ([]byte, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return nil, err
	}
	//open privatekeyfile
	buf, err := readKeyFile(privatekeypath)
	//pem decode
	privatekeyblock, _ := pem.Decode(buf)
	//X509 decode
	privateKey, err := x509.ParsePKCS1PrivateKey(privatekeyblock.Bytes)
	if err != nil {
		return nil, err
	}
	//decrypt the cipher
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertextBytes)

	return plainText, err
}

func SignInPostCommon(ctx *context.Context, form auth.SignInForm) {
	log.Info("start to login by ui=" + string(form.UserName))
	ctx.Data["Title"] = ctx.Tr("sign_in")
	orderedOAuth2Names, oauth2Providers, err := models.GetActiveOAuth2Providers()
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	ctx.Data["OrderedOAuth2Names"] = orderedOAuth2Names
	ctx.Data["OAuth2Providers"] = oauth2Providers
	ctx.Data["Title"] = ctx.Tr("sign_in")

	ctx.Data["PageIsSignIn"] = true

	ctx.Data["IsCourse"] = ctx.QueryBool("course")
	ctx.Data["EnableSSPI"] = models.IsSSPIEnabled()
	ctx.Data["EnableCloudBrain"] = true
	setRSAContext(ctx)
	if ctx.HasError() {
		ctx.HTML(200, tplSignIn)
		return
	}
	if setting.LoginEncrypt {
		decodePassword, err := RsaDecrypt(form.Password, setting.CustomPath+"/conf/privatekey.pem")
		if err == nil {
			log.Info("RSA decrypt succeed.")
			form.Password = string(decodePassword)
		} else {
			ctx.ServerError("UserSignIn", err)
			return
		}
	}
	u, err := models.UserSignIn(form.UserName, form.Password)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.RenderWithErr(ctx.Tr("form.username_password_incorrect"), tplSignIn, &form)
			log.Info("Failed authentication attempt for %s from %s", form.UserName, ctx.RemoteAddr())
		} else if models.IsErrEmailAlreadyUsed(err) {
			ctx.RenderWithErr(ctx.Tr("form.email_been_used"), tplSignIn, &form)
			log.Info("Failed authentication attempt for %s from %s", form.UserName, ctx.RemoteAddr())
		} else if models.IsErrUserProhibitLogin(err) {
			log.Info("Failed authentication attempt for %s from %s", form.UserName, ctx.RemoteAddr())
			ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
			ctx.HTML(200, "user/auth/prohibit_login")
		} else if models.IsErrUserInactive(err) {
			if setting.Service.RegisterEmailConfirm {
				ctx.Data["Title"] = ctx.Tr("auth.active_your_account")
				ctx.HTML(200, TplActivate)
			} else {
				log.Info("Failed authentication attempt for %s from %s", form.UserName, ctx.RemoteAddr())
				ctx.Data["Title"] = ctx.Tr("auth.prohibit_login")
				ctx.HTML(200, "user/auth/prohibit_login")
			}
		} else {
			log.Warn("UserSignIn", err)
			ctx.RenderWithErr(ctx.Tr("form.username_password_incorrect"), tplSignIn, &form)
			log.Info("Failed authentication attempt for %s from %s", form.UserName, ctx.RemoteAddr())
		}
		return
	}
	models.SaveLoginInfoToDb(ctx.Req.Request, u)
	// If this user is enrolled in 2FA, we can't sign the user in just yet.
	// Instead, redirect them to the 2FA authentication page.
	_, err = models.GetTwoFactorByUID(u.ID)
	if err != nil {
		if models.IsErrTwoFactorNotEnrolled(err) {
			handleSignIn(ctx, u, form.Remember)
		} else {
			ctx.ServerError("UserSignIn", err)
		}
		return
	}
	// User needs to use 2FA, save data and redirect to 2FA page.
	if err := ctx.Session.Set("twofaUid", u.ID); err != nil {
		ctx.ServerError("UserSignIn: Unable to set twofaUid in session", err)
		return
	}
	if err := ctx.Session.Set("twofaRemember", form.Remember); err != nil {
		ctx.ServerError("UserSignIn: Unable to set twofaRemember in session", err)
		return
	}
	if err := ctx.Session.Release(); err != nil {
		ctx.ServerError("UserSignIn: Unable to save session", err)
		return
	}
	regs, err := models.GetU2FRegistrationsByUID(u.ID)
	if err == nil && len(regs) > 0 {
		ctx.Redirect(setting.AppSubURL + "/user/u2f")
		return
	}

	ctx.Redirect(setting.AppSubURL + "/user/two_factor")
}

func SignInCloudBrainPost(ctx *context.Context, form auth.SignInForm) {
	ctx.Data["PageIsCloudBrainLogin"] = true
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/login/cloud_brain"
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()
	SignInPostCommon(ctx, form)
}

// SignInPost response for sign in request
func SignInPost(ctx *context.Context, form auth.SignInForm) {
	ctx.Data["PageIsLogin"] = true
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/login"
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()
	SignInPostCommon(ctx, form)
}

// TwoFactor shows the user a two-factor authentication page.
func TwoFactor(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("twofa")

	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}

	// Ensure user is in a 2FA session.
	if ctx.Session.Get("twofaUid") == nil {
		ctx.ServerError("UserSignIn", errors.New("not in 2FA session"))
		return
	}

	ctx.HTML(200, tplTwofa)
}

// TwoFactorPost validates a user's two-factor authentication token.
func TwoFactorPost(ctx *context.Context, form auth.TwoFactorAuthForm) {
	ctx.Data["Title"] = ctx.Tr("twofa")

	// Ensure user is in a 2FA session.
	idSess := ctx.Session.Get("twofaUid")
	if idSess == nil {
		ctx.ServerError("UserSignIn", errors.New("not in 2FA session"))
		return
	}

	id := idSess.(int64)
	twofa, err := models.GetTwoFactorByUID(id)
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}

	// Validate the passcode with the stored TOTP secret.
	ok, err := twofa.ValidateTOTP(form.Passcode)
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}

	if ok && twofa.LastUsedPasscode != form.Passcode {
		remember := ctx.Session.Get("twofaRemember").(bool)
		u, err := models.GetUserByID(id)
		if err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}

		if ctx.Session.Get("linkAccount") != nil {
			gothUser := ctx.Session.Get("linkAccountGothUser")
			if gothUser == nil {
				ctx.ServerError("UserSignIn", errors.New("not in LinkAccount session"))
				return
			}

			err = externalaccount.LinkAccountToUser(u, gothUser.(goth.User))
			if err != nil {
				ctx.ServerError("UserSignIn", err)
				return
			}
		}

		twofa.LastUsedPasscode = form.Passcode
		if err = models.UpdateTwoFactor(twofa); err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}

		handleSignIn(ctx, u, remember)
		return
	}

	ctx.RenderWithErr(ctx.Tr("auth.twofa_passcode_incorrect"), tplTwofa, auth.TwoFactorAuthForm{})
}

// TwoFactorScratch shows the scratch code form for two-factor authentication.
func TwoFactorScratch(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("twofa_scratch")

	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}

	// Ensure user is in a 2FA session.
	if ctx.Session.Get("twofaUid") == nil {
		ctx.ServerError("UserSignIn", errors.New("not in 2FA session"))
		return
	}

	ctx.HTML(200, tplTwofaScratch)
}

// TwoFactorScratchPost validates and invalidates a user's two-factor scratch token.
func TwoFactorScratchPost(ctx *context.Context, form auth.TwoFactorScratchAuthForm) {
	ctx.Data["Title"] = ctx.Tr("twofa_scratch")

	// Ensure user is in a 2FA session.
	idSess := ctx.Session.Get("twofaUid")
	if idSess == nil {
		ctx.ServerError("UserSignIn", errors.New("not in 2FA session"))
		return
	}

	id := idSess.(int64)
	twofa, err := models.GetTwoFactorByUID(id)
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}

	// Validate the passcode with the stored TOTP secret.
	if twofa.VerifyScratchToken(form.Token) {
		// Invalidate the scratch token.
		_, err = twofa.GenerateScratchToken()
		if err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}
		if err = models.UpdateTwoFactor(twofa); err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}

		remember := ctx.Session.Get("twofaRemember").(bool)
		u, err := models.GetUserByID(id)
		if err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}

		handleSignInFull(ctx, u, remember, false)
		ctx.Flash.Info(ctx.Tr("auth.twofa_scratch_used"))
		ctx.Redirect(setting.AppSubURL + "/user/settings/security")
		return
	}

	ctx.RenderWithErr(ctx.Tr("auth.twofa_scratch_token_incorrect"), tplTwofaScratch, auth.TwoFactorScratchAuthForm{})
}

// U2F shows the U2F login page
func U2F(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("twofa")
	ctx.Data["RequireU2F"] = true
	// Check auto-login.
	if checkAutoLogin(ctx) {
		return
	}

	// Ensure user is in a 2FA session.
	if ctx.Session.Get("twofaUid") == nil {
		ctx.ServerError("UserSignIn", errors.New("not in U2F session"))
		return
	}

	ctx.HTML(200, tplU2F)
}

// U2FChallenge submits a sign challenge to the browser
func U2FChallenge(ctx *context.Context) {
	// Ensure user is in a U2F session.
	idSess := ctx.Session.Get("twofaUid")
	if idSess == nil {
		ctx.ServerError("UserSignIn", errors.New("not in U2F session"))
		return
	}
	id := idSess.(int64)
	regs, err := models.GetU2FRegistrationsByUID(id)
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	if len(regs) == 0 {
		ctx.ServerError("UserSignIn", errors.New("no device registered"))
		return
	}
	challenge, err := u2f.NewChallenge(setting.U2F.AppID, setting.U2F.TrustedFacets)
	if err != nil {
		ctx.ServerError("u2f.NewChallenge", err)
		return
	}
	if err := ctx.Session.Set("u2fChallenge", challenge); err != nil {
		ctx.ServerError("UserSignIn: unable to set u2fChallenge in session", err)
		return
	}
	if err := ctx.Session.Release(); err != nil {
		ctx.ServerError("UserSignIn: unable to store session", err)
	}

	ctx.JSON(200, challenge.SignRequest(regs.ToRegistrations()))
}

// U2FSign authenticates the user by signResp
func U2FSign(ctx *context.Context, signResp u2f.SignResponse) {
	challSess := ctx.Session.Get("u2fChallenge")
	idSess := ctx.Session.Get("twofaUid")
	if challSess == nil || idSess == nil {
		ctx.ServerError("UserSignIn", errors.New("not in U2F session"))
		return
	}
	challenge := challSess.(*u2f.Challenge)
	id := idSess.(int64)
	regs, err := models.GetU2FRegistrationsByUID(id)
	if err != nil {
		ctx.ServerError("UserSignIn", err)
		return
	}
	for _, reg := range regs {
		r, err := reg.Parse()
		if err != nil {
			log.Fatal("parsing u2f registration: %v", err)
			continue
		}
		newCounter, authErr := r.Authenticate(signResp, *challenge, reg.Counter)
		if authErr == nil {
			reg.Counter = newCounter
			user, err := models.GetUserByID(id)
			if err != nil {
				ctx.ServerError("UserSignIn", err)
				return
			}
			remember := ctx.Session.Get("twofaRemember").(bool)
			if err := reg.UpdateCounter(); err != nil {
				ctx.ServerError("UserSignIn", err)
				return
			}

			if ctx.Session.Get("linkAccount") != nil {
				gothUser := ctx.Session.Get("linkAccountGothUser")
				if gothUser == nil {
					ctx.ServerError("UserSignIn", errors.New("not in LinkAccount session"))
					return
				}

				err = externalaccount.LinkAccountToUser(user, gothUser.(goth.User))
				if err != nil {
					ctx.ServerError("UserSignIn", err)
					return
				}
			}
			redirect := handleSignInFull(ctx, user, remember, false)
			if redirect == "" {
				redirect = setting.AppSubURL + "/"
			}
			ctx.PlainText(200, []byte(redirect))
			return
		}
	}
	ctx.Error(401)
}

// This handles the final part of the sign-in process of the user.
func handleSignIn(ctx *context.Context, u *models.User, remember bool) {
	handleSignInFull(ctx, u, remember, true)
}

func handleSignInFullNotRedirect(ctx *context.Context, u *models.User, remember bool, obeyRedirect bool) string {

	log.Info("enter here.")
	if remember {
		days := 86400 * setting.LogInRememberDays
		ctx.SetCookie(setting.CookieUserName, u.Name, days, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
		ctx.SetSuperSecureCookie(base.EncodeMD5(u.Rands+u.Passwd),
			setting.CookieRememberName, u.Name, days, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	}

	_ = ctx.Session.Delete("openid_verified_uri")
	_ = ctx.Session.Delete("openid_signin_remember")
	_ = ctx.Session.Delete("openid_determined_email")
	_ = ctx.Session.Delete("openid_determined_username")
	_ = ctx.Session.Delete("twofaUid")
	_ = ctx.Session.Delete("twofaRemember")
	_ = ctx.Session.Delete("u2fChallenge")
	_ = ctx.Session.Delete("linkAccount")
	if err := ctx.Session.Set("uid", u.ID); err != nil {
		log.Error("Error setting uid %d in session: %v", u.ID, err)
	}
	if err := ctx.Session.Set("uname", u.Name); err != nil {
		log.Error("Error setting uname %s session: %v", u.Name, err)
	}
	if err := ctx.Session.Release(); err != nil {
		log.Error("Unable to store session: %v", err)
	}

	// If the user does not have a locale set, we save the current one.
	if len(u.Language) == 0 {
		if len(ctx.GetCookie("lang")) != 0 {
			u.Language = ctx.GetCookie("lang")
		} else {
			u.Language = ctx.Locale.Language()
		}

		if err := models.UpdateUserCols(u, "language"); err != nil {
			log.Error(fmt.Sprintf("Error updating user language [user: %d, locale: %s]", u.ID, u.Language))
			return setting.AppSubURL + "/dashboard"
		}
	} else {
		// Language setting of the user use the one previously set
		if len(ctx.GetCookie("lang")) != 0 {
			u.Language = ctx.GetCookie("lang")
		}
	}

	ctx.SetCookie("lang", u.Language, nil, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	// Clear whatever CSRF has right now, force to generate a new one
	ctx.SetCookie(setting.CSRFCookieName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)

	// Register last login
	u.SetLastLogin()
	if err := models.UpdateUserCols(u, "last_login_unix"); err != nil {
		ctx.ServerError("UpdateUserCols", err)
		return setting.AppSubURL + "/dashboard"
	}

	return setting.AppSubURL + "/dashboard"
}

func handleSignInFull(ctx *context.Context, u *models.User, remember bool, obeyRedirect bool) string {
	if remember {
		days := 86400 * setting.LogInRememberDays
		ctx.SetCookie(setting.CookieUserName, u.Name, days, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
		ctx.SetSuperSecureCookie(base.EncodeMD5(u.Rands+u.Passwd),
			setting.CookieRememberName, u.Name, days, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	}

	_ = ctx.Session.Delete("openid_verified_uri")
	_ = ctx.Session.Delete("openid_signin_remember")
	_ = ctx.Session.Delete("openid_determined_email")
	_ = ctx.Session.Delete("openid_determined_username")
	_ = ctx.Session.Delete("twofaUid")
	_ = ctx.Session.Delete("twofaRemember")
	_ = ctx.Session.Delete("u2fChallenge")
	_ = ctx.Session.Delete("linkAccount")
	if err := ctx.Session.Set("uid", u.ID); err != nil {
		log.Error("Error setting uid %d in session: %v", u.ID, err)
	}
	if err := ctx.Session.Set("uname", u.Name); err != nil {
		log.Error("Error setting uname %s session: %v", u.Name, err)
	}
	if err := ctx.Session.Release(); err != nil {
		log.Error("Unable to store session: %v", err)
	}

	// If the user does not have a locale set, we save the current one.
	if len(u.Language) == 0 {
		if len(ctx.GetCookie("lang")) != 0 {
			u.Language = ctx.GetCookie("lang")
		} else {
			u.Language = ctx.Locale.Language()
		}

		if err := models.UpdateUserCols(u, "language"); err != nil {
			log.Error(fmt.Sprintf("Error updating user language [user: %d, locale: %s]", u.ID, u.Language))
			return setting.AppSubURL + "/dashboard"
		}
	} else {
		// Language setting of the user use the one previously set
		if len(ctx.GetCookie("lang")) != 0 {
			u.Language = ctx.GetCookie("lang")
		}
	}

	ctx.SetCookie("lang", u.Language, nil, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)

	// Clear whatever CSRF has right now, force to generate a new one
	ctx.SetCookie(setting.CSRFCookieName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)

	// Register last login
	u.SetLastLogin()
	if err := models.UpdateUserCols(u, "last_login_unix"); err != nil {
		ctx.ServerError("UpdateUserCols", err)
		return setting.AppSubURL + "/dashboard"
	}

	isCourse := ctx.QueryBool("course")
	if isCourse {
		redirectToCourse := setting.AppSubURL + "/" + setting.Course.OrgName
		ctx.RedirectToFirst(redirectToCourse)
		return redirectToCourse
	}
	if redirectTo := ctx.GetCookie("redirect_to"); len(redirectTo) > 0 && !util.IsExternalURL(redirectTo) {
		ctx.SetCookie("redirect_to", "", -1, setting.AppSubURL, "", setting.SessionConfig.Secure, true)
		if obeyRedirect {
			// 如果是web ide 需要判断是否是当前用户
			if strings.Contains(redirectTo, "/ide/project") {
				url, _ := url.Parse(setting.AppSubURL + redirectTo)
				query := url.Query()
				if query.Get("owner") == u.LowerName {
					ctx.RedirectToFirst(redirectTo)
				} else {
					ctx.Redirect(setting.AppSubURL + "/dashboard")
				}
			} else {
				ctx.RedirectToFirst(redirectTo)
			}
		}
		return redirectTo
	}
	if obeyRedirect {
		ctx.Redirect(setting.AppSubURL + "/dashboard")
	}
	return setting.AppSubURL + "/dashboard"
}

// SignInOAuth handles the OAuth2 login buttons
func SignInOAuth(ctx *context.Context) {
	provider := ctx.Params(":provider")

	loginSource, err := models.GetActiveOAuth2LoginSourceByName(provider)
	if err != nil {
		ctx.ServerError("SignIn", err)
		return
	}

	// try to do a direct callback flow, so we don't authenticate the user again but use the valid accesstoken to get the user
	user, gothUser, err := oAuth2UserLoginCallback(loginSource, ctx.Req.Request, ctx.Resp)
	if err == nil && user != nil {
		// we got the user without going through the whole OAuth2 authentication flow again
		handleOAuth2SignIn(user, gothUser, ctx, err)
		return
	}

	err = oauth2.Auth(loginSource.Name, ctx.Req.Request, ctx.Resp)
	if err != nil {
		ctx.ServerError("SignIn", err)
	}
	// redirect is done in oauth2.Auth
}

// SignInOAuthCallback handles the callback from the given provider
func SignInOAuthCallback(ctx *context.Context) {
	provider := ctx.Params(":provider")

	// first look if the provider is still active
	loginSource, err := models.GetActiveOAuth2LoginSourceByName(provider)
	if err != nil {
		ctx.ServerError("SignIn", err)
		return
	}

	if loginSource == nil {
		ctx.ServerError("SignIn", errors.New("No valid provider found, check configured callback url in provider"))
		return
	}

	u, gothUser, err := oAuth2UserLoginCallback(loginSource, ctx.Req.Request, ctx.Resp)

	handleOAuth2SignIn(u, gothUser, ctx, err)
}

func handleOAuth2SignIn(u *models.User, gothUser goth.User, ctx *context.Context, err error) {
	if err != nil {
		log.Error("Oauth callback err: %v", err)
		ctx.Redirect(setting.AppSubURL + "/user/login")
		return
	}

	if u == nil {

		// no existing user is found, request attach or new account
		if err := ctx.Session.Set("linkAccountGothUser", gothUser); err != nil {
			log.Error("Error setting linkAccountGothUser in session: %v", err)
		}
		if err := ctx.Session.Release(); err != nil {
			log.Error("Error storing session: %v", err)
		}

		if gothUser.Provider == setting.EduCoderConfig.Provider || gothUser.Provider == setting.GitlinkConfig.Provider || gothUser.Provider == setting.OsredmConfig.Provider {

			ctx.Redirect(setting.AppSubURL + "/user/link_account_new")

		} else {
			ctx.Redirect(setting.AppSubURL + "/user/link_account")
		}
		return
	}

	// If this user is enrolled in 2FA, we can't sign the user in just yet.
	// Instead, redirect them to the 2FA authentication page.
	_, err = models.GetTwoFactorByUID(u.ID)
	if err != nil {
		if !models.IsErrTwoFactorNotEnrolled(err) {
			ctx.ServerError("UserSignIn", err)
			return
		}

		if err := ctx.Session.Set("uid", u.ID); err != nil {
			log.Error("Error setting uid in session: %v", err)
		}
		if err := ctx.Session.Set("uname", u.Name); err != nil {
			log.Error("Error setting uname in session: %v", err)
		}
		if err := ctx.Session.Release(); err != nil {
			log.Error("Error storing session: %v", err)
		}

		// Clear whatever CSRF has right now, force to generate a new one
		ctx.SetCookie(setting.CSRFCookieName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)

		// Register last login
		u.SetLastLogin()
		if err := models.UpdateUserCols(u, "last_login_unix"); err != nil {
			ctx.ServerError("UpdateUserCols", err)
			return
		}

		// update external user information
		if err := models.UpdateExternalUser(u, gothUser); err != nil {
			log.Error("UpdateExternalUser failed: %v", err)
		}

		if redirectTo := ctx.GetCookie("redirect_to"); len(redirectTo) > 0 {
			ctx.SetCookie("redirect_to", "", -1, setting.AppSubURL, "", setting.SessionConfig.Secure, true)
			ctx.RedirectToFirst(redirectTo)
			return
		}
		if gothUser.Provider == setting.EduCoderConfig.Provider || gothUser.Provider == setting.GitlinkConfig.Provider || gothUser.Provider == setting.OsredmConfig.Provider {
			ctx.Redirect(setting.AppSubURL + "/dashboard")
		} else {
			ctx.Redirect(setting.AppSubURL + "/")
		}

		return
	}

	// User needs to use 2FA, save data and redirect to 2FA page.
	if err := ctx.Session.Set("twofaUid", u.ID); err != nil {
		log.Error("Error setting twofaUid in session: %v", err)
	}
	if err := ctx.Session.Set("twofaRemember", false); err != nil {
		log.Error("Error setting twofaRemember in session: %v", err)
	}
	if err := ctx.Session.Release(); err != nil {
		log.Error("Error storing session: %v", err)
	}

	// If U2F is enrolled -> Redirect to U2F instead
	regs, err := models.GetU2FRegistrationsByUID(u.ID)
	if err == nil && len(regs) > 0 {
		ctx.Redirect(setting.AppSubURL + "/user/u2f")
		return
	}

	ctx.Redirect(setting.AppSubURL + "/user/two_factor")
}

// OAuth2UserLoginCallback attempts to handle the callback from the OAuth2 provider and if successful
// login the user
func oAuth2UserLoginCallback(loginSource *models.LoginSource, request *http.Request, response http.ResponseWriter) (*models.User, goth.User, error) {
	gothUser, err := oauth2.ProviderCallback(loginSource.Name, request, response)

	if err != nil {
		if err.Error() == "securecookie: the value is too long" {
			log.Error("OAuth2 Provider %s returned too long a token. Current max: %d. Either increase the [OAuth2] MAX_TOKEN_LENGTH or reduce the information returned from the OAuth2 provider", loginSource.Name, setting.OAuth2.MaxTokenLength)
			err = fmt.Errorf("OAuth2 Provider %s returned too long a token. Current max: %d. Either increase the [OAuth2] MAX_TOKEN_LENGTH or reduce the information returned from the OAuth2 provider", loginSource.Name, setting.OAuth2.MaxTokenLength)
		}
		return nil, goth.User{}, err
	}

	user := &models.User{
		LoginName:   gothUser.UserID,
		LoginType:   models.LoginOAuth2,
		LoginSource: loginSource.ID,
	}

	hasUser, err := models.GetUser(user)
	if err != nil {
		return nil, goth.User{}, err
	}

	if hasUser {
		return user, gothUser, nil
	}

	// search in external linked users
	externalLoginUser := &models.ExternalLoginUser{
		ExternalID:    gothUser.UserID,
		LoginSourceID: loginSource.ID,
	}
	hasUser, err = models.GetExternalLogin(externalLoginUser)
	if err != nil {
		return nil, goth.User{}, err
	}
	if hasUser {
		user, err = models.GetUserByID(externalLoginUser.UserID)
		return user, gothUser, err
	}

	if gothUser.Provider == setting.EduCoderConfig.Provider || gothUser.Provider == setting.GitlinkConfig.Provider || gothUser.Provider == setting.OsredmConfig.Provider {
		if gothUser.Location != "" { //gothuser location save phone number
			userTemp := &models.User{
				PhoneNumber: gothUser.Location,
			}
			hasUser, err = models.GetUser(userTemp)
			if err != nil {
				return nil, goth.User{}, err
			}
			if hasUser {
				externalaccount.LinkAccountToUser(userTemp, gothUser)
				return userTemp, gothUser, nil
			} else {
				userTemp = &models.User{
					Name: gothUser.Name,
				}
				hasUser, err = models.GetUser(userTemp)
				if err != nil {
					return nil, goth.User{}, err
				}
				if !hasUser {
					userTemp = &models.User{
						Name:        gothUser.Name,
						Email:       "",
						Passwd:      "",
						IsActive:    !setting.Service.RegisterEmailConfirm,
						LoginType:   models.LoginOAuth2,
						LoginSource: loginSource.ID,
						LoginName:   gothUser.UserID,
						PhoneNumber: gothUser.Location,
						FullName:    gothUser.NickName,
						Language:    "zh-CN",
					}
					if err := models.CreateUser(userTemp); err != nil {
						return nil, goth.User{}, err
					} else {
						return userTemp, gothUser, nil
					}
				}

			}

		}
	}

	// no user found to login
	return nil, gothUser, nil

}

func LinkAccountNew(ctx *context.Context) {

	ctx.Data["Title"] = ctx.Tr("auth.oauth_signup_add_info_title")

	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/link_account_signup_new"

	gothUser := ctx.Session.Get("linkAccountGothUser")
	if gothUser == nil {
		ctx.ServerError("UserSignIn", errors.New("not in LinkAccount session"))
		return
	}

	uname := gothUser.(goth.User).Name
	ctx.Data["has_phone"] = false
	if gothUser.(goth.User).Location != "" {
		ctx.Data["has_phone"] = true
	}
	ctx.Data["user_name"] = uname
	if len(uname) != 0 {
		u, err := models.GetUserByName(uname)
		if err != nil && !models.IsErrUserNotExist(err) {
			ctx.ServerError("UserSignIn", err)
			return
		}
		if u != nil {
			ctx.Data["user_name"] = ""
		}
	}

	ctx.HTML(200, tplLinkAccountNew)
}

// LinkAccount shows the page where the user can decide to login or create a new account
func LinkAccount(ctx *context.Context) {
	ctx.Data["DisablePassword"] = !setting.Service.RequireExternalRegistrationPassword || setting.Service.AllowOnlyExternalRegistration
	ctx.Data["Title"] = ctx.Tr("link_account")
	ctx.Data["LinkAccountMode"] = true
	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha && setting.Service.RequireExternalRegistrationCaptcha
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["DisableRegistration"] = setting.Service.DisableRegistration
	ctx.Data["ShowRegistrationButton"] = false

	// use this to set the right link into the signIn and signUp templates in the link_account template
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/link_account_signin"
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/link_account_signup"

	gothUser := ctx.Session.Get("linkAccountGothUser")
	if gothUser == nil {
		ctx.ServerError("UserSignIn", errors.New("not in LinkAccount session"))
		return
	}

	uname := gothUser.(goth.User).NickName
	email := gothUser.(goth.User).Email
	ctx.Data["user_name"] = uname
	ctx.Data["email"] = email

	if len(email) != 0 {
		u, err := models.GetUserByEmail(email)
		if err != nil && !models.IsErrUserNotExist(err) {
			ctx.ServerError("UserSignIn", err)
			return
		}
		if u != nil {
			ctx.Data["user_exists"] = true
		}
	} else if len(uname) != 0 {
		u, err := models.GetUserByName(uname)
		if err != nil && !models.IsErrUserNotExist(err) {
			ctx.ServerError("UserSignIn", err)
			return
		}
		if u != nil {
			ctx.Data["user_exists"] = true
		}
	}

	ctx.HTML(200, tplLinkAccount)
}

// LinkAccountPostSignIn handle the coupling of external account with another account using signIn
func LinkAccountPostSignIn(ctx *context.Context, signInForm auth.SignInForm) {
	ctx.Data["DisablePassword"] = !setting.Service.RequireExternalRegistrationPassword || setting.Service.AllowOnlyExternalRegistration
	ctx.Data["Title"] = ctx.Tr("link_account")
	ctx.Data["LinkAccountMode"] = true
	ctx.Data["LinkAccountModeSignIn"] = true
	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha && setting.Service.RequireExternalRegistrationCaptcha
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["DisableRegistration"] = setting.Service.DisableRegistration
	ctx.Data["ShowRegistrationButton"] = false

	// use this to set the right link into the signIn and signUp templates in the link_account template
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/link_account_signin"
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/link_account_signup"

	gothUser := ctx.Session.Get("linkAccountGothUser")
	if gothUser == nil {
		ctx.ServerError("UserSignIn", errors.New("not in LinkAccount session"))
		return
	}

	if ctx.HasError() {
		ctx.HTML(200, tplLinkAccount)
		return
	}

	u, err := models.UserSignIn(signInForm.UserName, signInForm.Password)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.Data["user_exists"] = true
			ctx.RenderWithErr(ctx.Tr("form.username_password_incorrect"), tplLinkAccount, &signInForm)
		} else {
			ctx.ServerError("UserLinkAccount", err)
		}
		return
	}

	// If this user is enrolled in 2FA, we can't sign the user in just yet.
	// Instead, redirect them to the 2FA authentication page.
	_, err = models.GetTwoFactorByUID(u.ID)
	if err != nil {
		if !models.IsErrTwoFactorNotEnrolled(err) {
			ctx.ServerError("UserLinkAccount", err)
			return
		}

		err = externalaccount.LinkAccountToUser(u, gothUser.(goth.User))
		if err != nil {
			ctx.ServerError("UserLinkAccount", err)
			return
		}

		handleSignIn(ctx, u, signInForm.Remember)
		return
	}

	// User needs to use 2FA, save data and redirect to 2FA page.
	if err := ctx.Session.Set("twofaUid", u.ID); err != nil {
		log.Error("Error setting twofaUid in session: %v", err)
	}
	if err := ctx.Session.Set("twofaRemember", signInForm.Remember); err != nil {
		log.Error("Error setting twofaRemember in session: %v", err)
	}
	if err := ctx.Session.Set("linkAccount", true); err != nil {
		log.Error("Error setting linkAccount in session: %v", err)
	}
	if err := ctx.Session.Release(); err != nil {
		log.Error("Error storing session: %v", err)
	}

	// If U2F is enrolled -> Redirect to U2F instead
	regs, err := models.GetU2FRegistrationsByUID(u.ID)
	if err == nil && len(regs) > 0 {
		ctx.Redirect(setting.AppSubURL + "/user/u2f")
		return
	}

	ctx.Redirect(setting.AppSubURL + "/user/two_factor")
}

// LinkAccountPostRegister handle the creation of a new account for an external account using signUp
func LinkAccountPostRegister(ctx *context.Context, cpt *captcha.Captcha, form auth.RegisterForm) {
	// TODO Make insecure passwords optional for local accounts also,
	//      once email-based Second-Factor Auth is available
	ctx.Data["DisablePassword"] = !setting.Service.RequireExternalRegistrationPassword || setting.Service.AllowOnlyExternalRegistration
	ctx.Data["Title"] = ctx.Tr("link_account")
	ctx.Data["LinkAccountMode"] = true
	ctx.Data["LinkAccountModeRegister"] = true
	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha && setting.Service.RequireExternalRegistrationCaptcha
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["DisableRegistration"] = setting.Service.DisableRegistration
	ctx.Data["ShowRegistrationButton"] = false

	// use this to set the right link into the signIn and signUp templates in the link_account template
	ctx.Data["SignInLink"] = setting.AppSubURL + "/user/link_account_signin"
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/link_account_signup"

	gothUser := ctx.Session.Get("linkAccountGothUser")
	if gothUser == nil {
		ctx.ServerError("UserSignUp", errors.New("not in LinkAccount session"))
		return
	}

	if ctx.HasError() {
		ctx.HTML(200, tplLinkAccount)
		return
	}

	if setting.Service.DisableRegistration {
		ctx.Error(403)
		return
	}

	if setting.Service.EnableCaptcha && setting.Service.RequireExternalRegistrationCaptcha {
		var valid bool
		switch setting.Service.CaptchaType {
		case setting.ImageCaptcha:
			valid = cpt.VerifyReq(ctx.Req)
		case setting.ReCaptcha:
			valid, _ = recaptcha.Verify(form.GRecaptchaResponse)
		default:
			ctx.ServerError("Unknown Captcha Type", fmt.Errorf("Unknown Captcha Type: %s", setting.Service.CaptchaType))
			return
		}

		if !valid {
			ctx.Data["Err_Captcha"] = true
			ctx.RenderWithErr(ctx.Tr("form.captcha_incorrect"), tplLinkAccount, &form)
			return
		}
	}

	if setting.Service.AllowOnlyExternalRegistration || !setting.Service.RequireExternalRegistrationPassword {
		// In models.User an empty password is classed as not set, so we set form.Password to empty.
		// Eventually the database should be changed to indicate "Second Factor"-enabled accounts
		// (accounts that do not introduce the security vulnerabilities of a password).
		// If a user decides to circumvent second-factor security, and purposefully create a password,
		// they can still do so using the "Recover Account" option.
		form.Password = ""
	} else {
		if (len(strings.TrimSpace(form.Password)) > 0 || len(strings.TrimSpace(form.Retype)) > 0) && form.Password != form.Retype {
			ctx.Data["Err_Password"] = true
			ctx.RenderWithErr(ctx.Tr("form.password_not_match"), tplLinkAccount, &form)
			return
		}
		if len(strings.TrimSpace(form.Password)) > 0 && len(form.Password) < setting.MinPasswordLength {
			ctx.Data["Err_Password"] = true
			ctx.RenderWithErr(ctx.Tr("auth.password_too_short", setting.MinPasswordLength), tplLinkAccount, &form)
			return
		}
	}

	loginSource, err := models.GetActiveOAuth2LoginSourceByName(gothUser.(goth.User).Provider)
	if err != nil {
		ctx.ServerError("CreateUser", err)
	}

	u := &models.User{
		Name: form.UserName,
		//Email:       form.Email,
		Passwd:      form.Password,
		IsActive:    !setting.Service.RegisterEmailConfirm,
		LoginType:   models.LoginOAuth2,
		LoginSource: loginSource.ID,
		LoginName:   gothUser.(goth.User).UserID,
	}

	//nolint: dupl
	if err := models.CreateUser(u); err != nil {
		switch {
		case models.IsErrUserAlreadyExist(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("form.username_been_taken"), tplLinkAccount, &form)
		case models.IsErrEmailAlreadyUsed(err):
			ctx.Data["Err_Email"] = true
			ctx.RenderWithErr(ctx.Tr("form.email_been_used"), tplLinkAccount, &form)
		case models.IsErrNameReserved(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_reserved", err.(models.ErrNameReserved).Name), tplLinkAccount, &form)
		case models.IsErrNamePatternNotAllowed(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_pattern_not_allowed", err.(models.ErrNamePatternNotAllowed).Pattern), tplLinkAccount, &form)
		case models.IsErrNameCharsNotAllowed(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_chars_not_allowed", err.(models.ErrNameCharsNotAllowed).Name), tplLinkAccount, &form)
		default:
			ctx.ServerError("CreateUser", err)
		}
		return
	}
	log.Trace("Account created: %s", u.Name)

	// Auto-set admin for the only user.
	if models.CountUsers() == 1 {
		u.IsAdmin = true
		u.IsActive = true
		u.SetLastLogin()
		if err := models.UpdateUserCols(u, "is_admin", "is_active", "last_login_unix"); err != nil {
			ctx.ServerError("UpdateUser", err)
			return
		}
	}

	// update external user information
	if err := models.UpdateExternalUser(u, gothUser.(goth.User)); err != nil {
		log.Error("UpdateExternalUser failed: %v", err)
	}

	// Send confirmation email
	if setting.Service.RegisterEmailConfirm && u.ID > 1 {
		mailer.SendActivateAccountMail(ctx.Locale, u)

		ctx.Data["IsSendRegisterMail"] = true
		ctx.Data["Email"] = u.Email
		ctx.Data["ActiveCodeLives"] = timeutil.MinutesToFriendly(setting.Service.ActiveCodeLives, ctx.Locale.Language())
		ctx.HTML(200, TplActivate)

		if err := ctx.Cache.Put("MailResendLimit_"+u.LowerName, u.LowerName, 180); err != nil {
			log.Error("Set cache(MailResendLimit) fail: %v", err)
		}
		return
	}

	ctx.Redirect(setting.AppSubURL + "/user/login")
}

func LinkAccountPostRegisterNew(ctx *context.Context, form auth.RegisterOauthForm) {

	ctx.Data["Title"] = ctx.Tr("auth.oauth_signup_add_info_title")
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/link_account_signup_new"

	if !form.Agree {
		ctx.RenderWithErr(ctx.Tr("sign_up_agree_tips"), tplLinkAccountNew, &form)
		return
	}

	gothUser := ctx.Session.Get("linkAccountGothUser")
	if gothUser == nil {
		ctx.ServerError("UserSignUp", errors.New("not in LinkAccount session"))
		return
	}

	ctx.Data["has_phone"] = false
	if gothUser.(goth.User).Location != "" {
		ctx.Data["has_phone"] = true
	}

	if ctx.HasError() {
		ctx.HTML(200, tplLinkAccountNew)
		return
	}
	phoneNumber := strings.TrimSpace(form.PhoneNumber)
	if setting.PhoneService.Enabled && gothUser.(goth.User).Location == "" {
		verifyCode := strings.TrimSpace(form.VerifyCode)
		if !phoneService.IsVerifyCodeRight(phoneNumber, verifyCode) {
			ctx.RenderWithErr(ctx.Tr("phone.verify_code_fail"), tplLinkAccountNew, &form)
			return
		}
	}

	if gothUser.(goth.User).Location != "" {
		phoneNumber = strings.TrimSpace(gothUser.(goth.User).Location)
	}

	loginSource, err := models.GetActiveOAuth2LoginSourceByName(gothUser.(goth.User).Provider)
	if err != nil {
		ctx.ServerError("CreateUser", err)
	}

	u := &models.User{
		Name:        form.UserName,
		Email:       "",
		Passwd:      "",
		IsActive:    !setting.Service.RegisterEmailConfirm,
		LoginType:   models.LoginOAuth2,
		LoginSource: loginSource.ID,
		LoginName:   gothUser.(goth.User).UserID,
		Language:    "zh-CN",
		PhoneNumber: phoneNumber,
	}

	//nolint: dupl
	if err := models.CreateUser(u); err != nil {
		switch {
		case models.IsErrUserAlreadyExist(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("form.username_been_taken"), tplLinkAccountNew, &form)
		case models.IsErrEmailAlreadyUsed(err):
			ctx.Data["Err_Email"] = true
			ctx.RenderWithErr(ctx.Tr("form.email_been_used"), tplLinkAccountNew, &form)
		case models.IsErrNameReserved(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_reserved", err.(models.ErrNameReserved).Name), tplLinkAccountNew, &form)
		case models.IsErrNamePatternNotAllowed(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_pattern_not_allowed", err.(models.ErrNamePatternNotAllowed).Pattern), tplLinkAccountNew, &form)
		case models.IsErrNameCharsNotAllowed(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_chars_not_allowed", err.(models.ErrNameCharsNotAllowed).Name), tplLinkAccountNew, &form)
		default:
			ctx.ServerError("CreateUser", err)
		}
		return
	}
	log.Trace("Account created: %s", u.Name)

	// Auto-set admin for the only user.
	if models.CountUsers() == 1 {
		u.IsAdmin = true
		u.IsActive = true
		u.SetLastLogin()
		if err := models.UpdateUserCols(u, "is_admin", "is_active", "last_login_unix"); err != nil {
			ctx.ServerError("UpdateUser", err)
			return
		}
	}

	// update external user information
	if err := models.UpdateExternalUser(u, gothUser.(goth.User)); err != nil {
		log.Error("UpdateExternalUser failed: %v", err)
	}

	handleOAuth2SignIn(u, gothUser.(goth.User), ctx, nil)

}

// HandleSignOut resets the session and sets the cookies
func HandleSignOut(ctx *context.Context) {
	_ = ctx.Session.Flush()
	_ = ctx.Session.Destroy(ctx.Context)
	ctx.SetCookie(setting.CookieUserName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	ctx.SetCookie(setting.CookieRememberName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	ctx.SetCookie(setting.CSRFCookieName, "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true)
	ctx.SetCookie("lang", "", -1, setting.AppSubURL, setting.SessionConfig.Domain, setting.SessionConfig.Secure, true) // Setting the lang cookie will trigger the middleware to reset the language ot previous state.
	ctx.SetCookie("redirect_to", "", -1, setting.AppSubURL)                                                            // logout default should set redirect to to default
}

// SignOut sign out from login status
func SignOut(ctx *context.Context) {
	if ctx.User != nil {
		eventsource.GetManager().SendMessageBlocking(ctx.User.ID, &eventsource.Event{
			Name: "logout",
			Data: ctx.Session.ID(),
		})

		go SignOutNotify(ctx.User.ID)
	}
	HandleSignOut(ctx)
	ctx.Redirect(setting.AppSubURL + "/")
}

func SignOutNotify(uid int64) {
	NotifyPipelineUserLogout(uid)
}

func SignUpScanCode(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_up")
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/sign_up"

	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["PageIsSignUp"] = true

	//Show Disabled Registration message if DisableRegistration or AllowOnlyExternalRegistration options are true
	ctx.Data["DisableRegistration"] = setting.Service.DisableRegistration || setting.Service.AllowOnlyExternalRegistration
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()
	setRSAContext(ctx)

	ctx.HTML(200, tplSignUpScanCode)
}

// SignUp render the register page
func SignUp(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("sign_up")

	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/sign_up"

	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["PageIsSignUp"] = true

	//Show Disabled Registration message if DisableRegistration or AllowOnlyExternalRegistration options are true
	ctx.Data["DisableRegistration"] = setting.Service.DisableRegistration || setting.Service.AllowOnlyExternalRegistration
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()
	setRSAContext(ctx)

	ctx.HTML(200, tplSignUp)
}

// SignUpPost response for sign up information submission
func SignUpPost(ctx *context.Context, cpt *captcha.Captcha, form auth.RegisterForm) {
	ctx.Data["Title"] = ctx.Tr("sign_up")
	invitationCode := ctx.Query("invitation_code")
	ctx.Data["SignUpLink"] = setting.AppSubURL + "/user/sign_up"
	ctx.Data["invitationCode"] = invitationCode
	ctx.Data["EnableCaptcha"] = setting.Service.EnableCaptcha
	ctx.Data["RecaptchaURL"] = setting.Service.RecaptchaURL
	ctx.Data["CaptchaType"] = setting.Service.CaptchaType
	ctx.Data["RecaptchaSitekey"] = setting.Service.RecaptchaSitekey
	ctx.Data["PageIsSignUp"] = true
	ctx.Data["WechatSignUpRequired"] = setting.WechatSignUpRequired
	ctx.Data["ActivityTpl"] = getActivityTpl()
	setRSAContext(ctx)

	if setting.LoginEncrypt {
		decodePassword, err := RsaDecrypt(form.Password, setting.CustomPath+"/conf/privatekey.pem")
		if err == nil {
			log.Info("RSA decrypt succeed.")
			form.Password = string(decodePassword)
		} else {
			ctx.ServerError("UserSignUp", err)
			return
		}
		decodeRetype, err := RsaDecrypt(form.Retype, setting.CustomPath+"/conf/privatekey.pem")
		if err == nil {
			log.Info("RSA decrypt succeed.")
			form.Retype = string(decodeRetype)
		} else {
			ctx.ServerError("UserSignUp", err)
			return
		}
	}

	//Permission denied if DisableRegistration or AllowOnlyExternalRegistration options are true
	if setting.Service.DisableRegistration || setting.Service.AllowOnlyExternalRegistration {
		ctx.Error(403)
		return
	}

	if ctx.HasError() {
		ctx.HTML(200, tplSignUp)
		return
	}

	if !form.Agree {
		ctx.RenderWithErr(ctx.Tr("sign_up_agree_tips"), tplSignUp, &form)
		return
	}

	//校验是否已扫微信
	var wechatOpenId string
	if setting.WechatSignUpRequired {
		if form.SignUpId == "" {
			ctx.Data["Err_Wechat"] = true
			ctx.Data["Err_Wechat_Type"] = "should_bind_wechat"
			ctx.Data["Err_Wechat_Content"] = ctx.Tr("form.should_bind_wechat")
			ctx.RenderWithErr("", tplSignUp, &form)
			return
		}
		wechatOpenId = sign_up_service.GetWechatOpenIdBySignUpId(form.SignUpId)
		if wechatOpenId == "" {
			ctx.Data["Err_Wechat"] = true
			ctx.Data["Err_Wechat_Type"] = "qr_code_expire"
			ctx.Data["Err_Wechat_Content"] = ctx.Tr("form.qr_code_expire")
			ctx.RenderWithErr("", tplSignUp, &form)
			return
		}
		if !sign_up_service.IsWechatOpenIdAvailable(wechatOpenId) {
			ctx.Data["Err_Wechat"] = true
			ctx.Data["Err_Wechat_Type"] = "wechat_account_used"
			ctx.Data["Err_Wechat_Content"] = ctx.Tr("form.wechat_account_used")
			ctx.RenderWithErr("", tplSignUp, &form)
			return
		}
	}

	if setting.Service.EnableCaptcha {
		var valid bool
		switch setting.Service.CaptchaType {
		case setting.ImageCaptcha:
			valid = cpt.VerifyReq(ctx.Req)
		case setting.ReCaptcha:
			valid, _ = recaptcha.Verify(form.GRecaptchaResponse)
		default:
			ctx.ServerError("Unknown Captcha Type", fmt.Errorf("Unknown Captcha Type: %s", setting.Service.CaptchaType))
			return
		}

		if !valid {
			ctx.Data["Err_Captcha"] = true
			ctx.RenderWithErr(ctx.Tr("form.captcha_incorrect"), tplSignUp, &form)
			return
		}
	}

	if strings.ToLower(invitationCode) == strings.ToLower(form.UserName) {
		ctx.RenderWithErr(ctx.Tr("user.form.username_and_invited_code_duplicated"), tplSignUp, &form)
		return
	}

	//if !form.IsEmailDomainWhitelisted() {
	//	ctx.RenderWithErr(ctx.Tr("auth.email_domain_blacklisted"), tplSignUp, &form)
	//	return
	//}
	if form.Password != form.Retype {
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(ctx.Tr("form.password_not_match"), tplSignUp, &form)
		return
	}
	if len(form.Password) < setting.MinPasswordLength {
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(ctx.Tr("auth.password_too_short", setting.MinPasswordLength), tplSignUp, &form)
		return
	}
	if !password.IsComplexEnough(form.Password) {
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(password.BuildComplexityError(ctx), tplSignUp, &form)
		return
	}

	if setting.PhoneService.Enabled {
		phoneNumber := strings.TrimSpace(form.PhoneNumber)
		verifyCode := strings.TrimSpace(form.VerifyCode)
		if !phoneService.IsVerifyCodeRight(phoneNumber, verifyCode) {
			ctx.RenderWithErr(ctx.Tr("phone.verify_code_fail"), tplSignUp, &form)
			return
		}
	}

	if form.Occupation == 0 {
		ctx.Data["Err_Occupation"] = true
		ctx.RenderWithErr(ctx.Tr("occupation_empty"), tplSignUp, &form)
		return
	}

	u := &models.User{
		Name: form.UserName,
		//Email:        form.Email,
		Passwd:      form.Password,
		PhoneNumber: strings.TrimSpace(form.PhoneNumber),
		IsActive:    !setting.Service.RegisterEmailConfirm,
		Occupation:  form.Occupation,
	}
	if err := models.CreateUser(u); err != nil {
		switch {
		case models.IsErrUserAlreadyExist(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("form.username_been_taken"), tplSignUp, &form)
		case models.IsErrEmailAlreadyUsed(err):
			ctx.Data["Err_Email"] = true
			ctx.RenderWithErr(ctx.Tr("form.email_been_used"), tplSignUp, &form)
		case models.IsErrNameReserved(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_reserved", err.(models.ErrNameReserved).Name), tplSignUp, &form)
		case models.IsErrNamePatternNotAllowed(err):
			ctx.Data["Err_UserName"] = true
			ctx.RenderWithErr(ctx.Tr("user.form.name_pattern_not_allowed", err.(models.ErrNamePatternNotAllowed).Pattern), tplSignUp, &form)
		default:
			ctx.ServerError("CreateUser", err)
		}
		return
	}

	if setting.WechatSignUpRequired {
		if err := wechat.BindWechat(u.ID, wechatOpenId); err == nil {
			go notification.NotifyWechatBind(u, wechatOpenId)
		}
	}
	log.Trace("Account created: %s", u.Name, ctx.Data["MsgID"])

	log.Info("enter here, and form.InvitaionCode =" + invitationCode)
	if invitationCode != "" {
		tmpErr := RegisteUserByInvitaionCode(invitationCode, u)
		if tmpErr != nil {
			log.Error("RegisteUserByInvitaionCode err.u=%+v invitationCode=%s", u, invitationCode)
		}
	}

	// Auto-set admin for the only user.
	if models.CountUsers() == 1 {
		u.IsAdmin = true
		u.IsActive = true
		u.SetLastLogin()
		if err := models.UpdateUserCols(u, "is_admin", "is_active", "last_login_unix"); err != nil {
			ctx.ServerError("UpdateUser", err)
			return
		}
	}
	if setting.EnableAIAvatar {

		if err := u.GenerateRandomAvatarUseDiffustionModel(); err != nil {
			log.Error("GenerateRandomAvatarUseDiffustionModel: %v", err)
		}

	}

	ctx.Flash.Success(ctx.Tr("auth.sign_up_successful"))
	handleSignInFull(ctx, u, false, true)
}

// update user emailAddress
func UpdateEmailPost(ctx *context.Context, form auth.UpdateEmailForm) {
	newEmailAddress := ctx.Query("NewEmail")
	if newEmailAddress == "" {
		log.Error("please input the newEmail")
		return
	}
	if used, _ := models.IsEmailUsed(newEmailAddress); used {
		ctx.RenderWithErr(ctx.Tr("form.email_been_used"), TplActivate, &form)
		return
	}
	user := ctx.User
	email, err := models.GetEmailAddressByIDAndEmail(user.ID, user.Email)
	if err != nil {
		ctx.ServerError("GetEmailAddressByIDAndEmail failed", err)
		return
	}
	err = email.UpdateEmailAddress(newEmailAddress)
	if err != nil {
		ctx.ServerError("UpdateEmailAddress failed", err)
		return
	}
	ctx.Data["SignedUser.Email"] = newEmailAddress
	ctx.User.Email = newEmailAddress
	Activate(ctx)

}

// Activate render activate user page
func Activate(ctx *context.Context) {
	code := ctx.Query("code")
	if len(code) == 0 {
		ctx.Data["IsActivatePage"] = true
		if ctx.User.IsActive {
			ctx.Error(404)
			return
		}
		// Resend confirmation email.
		if setting.Service.RegisterEmailConfirm {
			if ctx.Cache.IsExist("MailResendLimit_" + ctx.User.LowerName) {
				ctx.Data["ResendLimited"] = true
			} else {
				ctx.Data["ActiveCodeLives"] = timeutil.MinutesToFriendly(setting.Service.ActiveCodeLives, ctx.Locale.Language())
				mailer.SendActivateAccountMail(ctx.Locale, ctx.User)

				if err := ctx.Cache.Put("MailResendLimit_"+ctx.User.LowerName, ctx.User.LowerName, 180); err != nil {
					log.Error("Set cache(MailResendLimit) fail: %v", err)
				}
			}
		} else {
			ctx.Data["ServiceNotEnabled"] = true
		}
		ctx.HTML(200, TplActivate)
		return
	}

	// Verify code.
	if user := models.VerifyUserActiveCode(code); user != nil {
		user.IsActive = true
		var err error
		if user.Rands, err = models.GetUserSalt(); err != nil {
			ctx.ServerError("UpdateUser", err)
			return
		}
		if err := models.UpdateUserCols(user, "is_active", "rands"); err != nil {
			if models.IsErrUserNotExist(err) {
				ctx.Error(404)
			} else {
				ctx.ServerError("UpdateUser", err)
			}
			return
		}

		log.Trace("User activated: %s", user.Name)

		if err := ctx.Session.Set("uid", user.ID); err != nil {
			log.Error(fmt.Sprintf("Error setting uid in session: %v", err))
		}
		if err := ctx.Session.Set("uname", user.Name); err != nil {
			log.Error(fmt.Sprintf("Error setting uname in session: %v", err))
		}
		if err := ctx.Session.Release(); err != nil {
			log.Error("Error storing session: %v", err)
		}

		email, err := models.GetEmailAddressByIDAndEmail(user.ID, user.Email)
		if err != nil || email == nil {
			log.Error("GetEmailAddressByIDAndEmail failed", ctx.Data["MsgID"])
		} else {
			if err := email.Activate(); err != nil {
				log.Error("Activate failed: %v", err, ctx.Data["MsgID"])
			}
		}

		ctx.Flash.Success(ctx.Tr("auth.account_activated"))
		ctx.Redirect(setting.AppSubURL + "/")
		return
	}

	ctx.Data["IsActivateFailed"] = true
	ctx.HTML(200, TplActivate)
}

// ActivateEmail render the activate email page
func ActivateEmail(ctx *context.Context) {
	code := ctx.Query("code")
	emailStr := ctx.Query("email")

	// Verify code.
	if email := models.VerifyActiveEmailCode(code, emailStr); email != nil {
		if err := email.Activate(); err != nil {
			ctx.ServerError("ActivateEmail", err)
		}

		log.Trace("Email activated: %s", email.Email)
		ctx.Flash.Success(ctx.Tr("settings.add_email_success"))

		if u, err := models.GetUserByID(email.UID); err != nil {
			log.Warn("GetUserByID: %d", email.UID)
		} else {
			// Allow user to validate more emails
			_ = ctx.Cache.Delete("MailResendLimit_" + u.LowerName)

			if u.Email == "" {
				//未设置主邮箱时则默认第一个激活的邮箱为主要邮箱
				if err := models.MakeEmailPrimary(email); err != nil {
					log.Error("MakeEmailPrimary default err. %v", err)
				}
			}
		}
	}

	// FIXME: e-mail verification does not require the user to be logged in,
	// so this could be redirecting to the login page.
	// Should users be logged in automatically here? (consider 2FA requirements, etc.)
	ctx.Redirect(setting.AppSubURL + "/user/settings/account")
}

// ForgotPasswd render the forget pasword page
func ForgotPasswd(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("auth.forgot_password_title")
	forgetType := ctx.Query("type")

	if forgetType == "phone" {
		if !setting.PhoneService.Enabled {
			ctx.Data["IsResetDisable"] = true
			ctx.HTML(200, tplForgotPasswordPhone)
			return
		}
		ctx.Data["IsResetRequest"] = true
		ctx.HTML(200, tplForgotPasswordPhone)
	} else {

		if setting.MailService == nil {
			ctx.Data["IsResetDisable"] = true
			ctx.HTML(200, tplForgotPassword)
			return
		}

		email := ctx.Query("email")
		ctx.Data["Email"] = email

		ctx.Data["IsResetRequest"] = true
		ctx.HTML(200, tplForgotPassword)
	}
}

// ForgotPasswdPost response for forget password request
func ForgotPasswdPost(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("auth.forgot_password_title")

	if setting.MailService == nil {
		ctx.NotFound("ForgotPasswdPost", nil)
		return
	}
	ctx.Data["IsResetRequest"] = true

	email := ctx.Query("email")
	ctx.Data["Email"] = email

	u, err := models.GetUserByMainEmail(email)
	if err != nil {
		if models.IsErrUserNotExist(err) {
			ctx.Data["ResetPwdCodeLives"] = timeutil.MinutesToFriendly(setting.Service.ResetPwdCodeLives, ctx.Locale.Language())
			ctx.Data["IsResetSent"] = false
			if used, _ := models.IsEmailUsed(email); used {
				ctx.RenderWithErr(ctx.Tr("auth.email_not_main"), tplForgotPassword, nil)
			} else {
				ctx.RenderWithErr(ctx.Tr("auth.email_not_right"), tplForgotPassword, nil)
			}
			return
		}

		ctx.ServerError("user.ResetPasswd(check existence)", err)
		return
	}

	if !u.IsLocal() && !u.IsOAuth2() && !u.IsCloudBrain() {
		ctx.Data["Err_Email"] = true
		ctx.RenderWithErr(ctx.Tr("auth.non_local_account"), tplForgotPassword, nil)
		return
	}

	if ctx.Cache.IsExist("MailResendLimit_" + u.LowerName) {
		ctx.Data["ResendLimited"] = true
		ctx.HTML(200, tplForgotPassword)
		return
	}

	mailer.SendResetPasswordMail(ctx.Locale, u)

	if err = ctx.Cache.Put("MailResendLimit_"+u.LowerName, u.LowerName, 180); err != nil {
		log.Error("Set cache(MailResendLimit) fail: %v", err)
	}

	ctx.Data["ResetPwdCodeLives"] = timeutil.MinutesToFriendly(setting.Service.ResetPwdCodeLives, ctx.Locale.Language())
	ctx.Data["IsResetSent"] = true
	ctx.HTML(200, tplForgotPassword)
}

func commonResetPassword(ctx *context.Context) (*models.User, *models.TwoFactor) {
	code := ctx.Query("code")

	ctx.Data["Title"] = ctx.Tr("auth.reset_password")
	ctx.Data["Code"] = code

	if nil != ctx.User {
		ctx.Data["user_signed_in"] = true
	}

	if len(code) == 0 {
		ctx.Flash.Error(ctx.Tr("auth.invalid_code"))
		return nil, nil
	}

	// Fail early, don't frustrate the user
	u := models.VerifyUserActiveCode(code)
	if u == nil {
		ctx.Flash.Error(ctx.Tr("auth.invalid_code"))
		return nil, nil
	}

	twofa, err := models.GetTwoFactorByUID(u.ID)
	if err != nil {
		if !models.IsErrTwoFactorNotEnrolled(err) {
			ctx.Error(http.StatusInternalServerError, "CommonResetPassword", err.Error())
			return nil, nil
		}
	} else {
		ctx.Data["has_two_factor"] = true
		ctx.Data["scratch_code"] = ctx.QueryBool("scratch_code")
	}

	// Show the user that they are affecting the account that they intended to
	ctx.Data["user_email"] = u.Email

	if nil != ctx.User && u.ID != ctx.User.ID {
		ctx.Flash.Error(ctx.Tr("auth.reset_password_wrong_user", ctx.User.Email, u.Email))
		return nil, nil
	}

	return u, twofa
}

// ResetPasswd render the account recovery page
func ResetPasswd(ctx *context.Context) {
	ctx.Data["IsResetForm"] = true

	commonResetPassword(ctx)
	if ctx.Written() {
		return
	}

	ctx.HTML(200, tplResetPassword)
}

// ResetPasswdPost response from account recovery request
func ResetPasswdPost(ctx *context.Context) {
	u, twofa := commonResetPassword(ctx)
	if ctx.Written() {
		return
	}

	if u == nil {
		// Flash error has been set
		ctx.HTML(200, tplResetPassword)
		return
	}

	// Validate password length.
	passwd := ctx.Query("password")
	if len(passwd) < setting.MinPasswordLength {
		ctx.Data["IsResetForm"] = true
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(ctx.Tr("auth.password_too_short", setting.MinPasswordLength), tplResetPassword, nil)
		return
	} else if !password.IsComplexEnough(passwd) {
		ctx.Data["IsResetForm"] = true
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(password.BuildComplexityError(ctx), tplResetPassword, nil)
		return
	}

	// Handle two-factor
	regenerateScratchToken := false
	if twofa != nil {
		if ctx.QueryBool("scratch_code") {
			if !twofa.VerifyScratchToken(ctx.Query("token")) {
				ctx.Data["IsResetForm"] = true
				ctx.Data["Err_Token"] = true
				ctx.RenderWithErr(ctx.Tr("auth.twofa_scratch_token_incorrect"), tplResetPassword, nil)
				return
			}
			regenerateScratchToken = true
		} else {
			passcode := ctx.Query("passcode")
			ok, err := twofa.ValidateTOTP(passcode)
			if err != nil {
				ctx.Error(http.StatusInternalServerError, "ValidateTOTP", err.Error())
				return
			}
			if !ok || twofa.LastUsedPasscode == passcode {
				ctx.Data["IsResetForm"] = true
				ctx.Data["Err_Passcode"] = true
				ctx.RenderWithErr(ctx.Tr("auth.twofa_passcode_incorrect"), tplResetPassword, nil)
				return
			}

			twofa.LastUsedPasscode = passcode
			if err = models.UpdateTwoFactor(twofa); err != nil {
				ctx.ServerError("ResetPasswdPost: UpdateTwoFactor", err)
				return
			}
		}
	}

	var err error
	if u.Rands, err = models.GetUserSalt(); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}
	if u.Salt, err = models.GetUserSalt(); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}
	u.HashPassword(passwd)
	u.MustChangePassword = false
	if u.LoginType == models.LoginCloudBrain {
		u.LoginType = models.LoginNoType
	}
	if err := models.UpdateUserCols(u, "must_change_password", "passwd", "rands", "salt", "login_type"); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}

	log.Trace("User password reset: %s", u.Name)
	ctx.Data["IsResetFailed"] = true
	remember := len(ctx.Query("remember")) != 0

	if regenerateScratchToken {
		// Invalidate the scratch token.
		_, err = twofa.GenerateScratchToken()
		if err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}
		if err = models.UpdateTwoFactor(twofa); err != nil {
			ctx.ServerError("UserSignIn", err)
			return
		}

		handleSignInFull(ctx, u, remember, false)
		ctx.Flash.Info(ctx.Tr("auth.twofa_scratch_used"))
		ctx.Redirect(setting.AppSubURL + "/user/settings/security")
		return
	}

	handleSignInFull(ctx, u, remember, true)
}

func ResetPasswdByPhonePost(ctx *context.Context, form auth.ResetPassWordByPhoneForm) {
	phoneNumber := strings.TrimSpace(form.PhoneNumber)
	verifyCode := strings.TrimSpace(form.VerifyCode)
	isRight := phoneService.IsVerifyCodeRight(phoneNumber, verifyCode)
	if !isRight {
		ctx.RenderWithErr(ctx.Tr("phone.verify_code_fail"), tplForgotPasswordPhone, form)
		return
	}

	passwd := strings.TrimSpace(form.Password)
	if len(passwd) < setting.MinPasswordLength {
		ctx.RenderWithErr(ctx.Tr("auth.password_too_short", setting.MinPasswordLength), tplForgotPasswordPhone, form)
		return
	} else if !password.IsComplexEnough(passwd) {
		ctx.RenderWithErr(password.BuildComplexityError(ctx), tplForgotPasswordPhone, form)
		return
	}

	u, err := models.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		log.Error("fail to query by phone number", err)
		ctx.RenderWithErr(ctx.Tr("phone.query_err", setting.MinPasswordLength), tplForgotPasswordPhone, form)
		return
	}

	if nil != ctx.User && u.ID != ctx.User.ID {
		ctx.RenderWithErr(ctx.Tr("auth.reset_password_wrong_user", ctx.User.Email, u.Email), tplForgotPasswordPhone, form)
		return
	}

	if u.Rands, err = models.GetUserSalt(); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}
	if u.Salt, err = models.GetUserSalt(); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}
	u.HashPassword(passwd)
	u.MustChangePassword = false
	if u.LoginType == models.LoginCloudBrain {
		u.LoginType = models.LoginNoType
	}
	if err := models.UpdateUserCols(u, "must_change_password", "passwd", "rands", "salt", "login_type"); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}

	handleSignInFull(ctx, u, form.Remember, true)

}

// MustChangePassword renders the page to change a user's password
func MustChangePassword(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
	ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/settings/change_password"
	ctx.HTML(200, tplMustChangePassword)
}

// MustChangePasswordPost response for updating a user's password after his/her
// account was created by an admin
func MustChangePasswordPost(ctx *context.Context, cpt *captcha.Captcha, form auth.MustChangePasswordForm) {
	ctx.Data["Title"] = ctx.Tr("auth.must_change_password")
	ctx.Data["ChangePasscodeLink"] = setting.AppSubURL + "/user/settings/change_password"
	if ctx.HasError() {
		ctx.HTML(200, tplMustChangePassword)
		return
	}
	u := ctx.User
	// Make sure only requests for users who are eligible to change their password via
	// this method passes through
	if !u.MustChangePassword {
		ctx.ServerError("MustUpdatePassword", errors.New("cannot update password.. Please visit the settings page"))
		return
	}

	if form.Password != form.Retype {
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(ctx.Tr("form.password_not_match"), tplMustChangePassword, &form)
		return
	}

	if len(form.Password) < setting.MinPasswordLength {
		ctx.Data["Err_Password"] = true
		ctx.RenderWithErr(ctx.Tr("auth.password_too_short", setting.MinPasswordLength), tplMustChangePassword, &form)
		return
	}

	var err error
	if u.Salt, err = models.GetUserSalt(); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}

	u.HashPassword(form.Password)
	u.MustChangePassword = false

	if err := models.UpdateUserCols(u, "must_change_password", "passwd", "salt"); err != nil {
		ctx.ServerError("UpdateUser", err)
		return
	}

	ctx.Flash.Success(ctx.Tr("settings.change_password_success"))

	log.Trace("User updated password: %s", u.Name)

	if redirectTo := ctx.GetCookie("redirect_to"); len(redirectTo) > 0 && !util.IsExternalURL(redirectTo) {
		ctx.SetCookie("redirect_to", "", -1, setting.AppSubURL)
		ctx.RedirectToFirst(redirectTo)
		return
	}

	ctx.Redirect(setting.AppSubURL + "/")
}

func CreateSlideImageInfo(ctx *context.Context, slideImage *slideimage.SlideImage) {
	id, _, _ := slideImage.CreateCode()
	ctx.JSON(http.StatusOK, models.BaseMessage{0, id})
}

func VerifySlideImage(ctx *context.Context, slideImage *slideimage.SlideImage, form auth.SlideImageForm) {
	if slideImage.Verify(form.SlideID, form.X) {
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
	} else {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(""))
	}
}

func VerifyGuestSlideImage(ctx *context.Context, slideImage *slideimage.SlideImage, form auth.SlideImageForm) {
	guest := auth.GetGuest()
	if guest.Enabled {
		ip := ctx.RemoteAddr()
		if setting.CustomRealUrlHeader != "" {
			addr := ctx.Req.Header.Get(setting.CustomRealUrlHeader)
			if len(addr) > 0 {
				ip = addr
			}
		}
		if slideImage.Verify(form.SlideID, form.X) {
			if guest.IsInTempList(ip) {
				guest.AddWhiteList(ip)
				guest.ResetMaxFailedCount(ip)
				ctx.JSON(http.StatusOK, models.BaseOKMessageApi)

			} else {
				ctx.JSON(http.StatusOK, models.BaseMessageApi{2, "time out"})
			}

		} else {
			if guest.IsInTempList(ip) {
				guest.AddMaxFailedCount(ip)
				if guest.GetMaxFailedCount(ip) > guest.MaxFailedCount {
					guest.AddBlackList(ip)
					guest.ResetMaxFailedCount(ip)
					ctx.JSON(http.StatusOK, models.BaseMessageApi{2, "time out"})
					return
				}
				ctx.JSON(http.StatusOK, models.BaseErrorMessageApi(""))
			} else {
				ctx.JSON(http.StatusOK, models.BaseMessageApi{2, "time out"})
			}
		}
	} else {
		ctx.JSON(http.StatusOK, models.BaseMessageApi{2, "not enabled"})
	}
}

func BindPhone(ctx *context.Context, form auth.PhoneNumberCodeForm) {
	if strings.TrimSpace(form.PhoneNumber) != "" && strings.TrimSpace(form.VerifyCode) != "" && phoneService.IsVerifyCodeRight(strings.TrimSpace(form.PhoneNumber), strings.TrimSpace(form.VerifyCode)) {

		ctx.User.PhoneNumber = strings.TrimSpace(form.PhoneNumber)
		if err := models.UpdateUserSetting(ctx.User); err != nil {
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.bind_phone_fail")))
			return
		}
		ctx.JSON(http.StatusOK, models.BaseOKMessage)
		return
	}

	ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.verify_code_fail")))

}

func SendVerifyCode(ctx *context.Context, slideImage *slideimage.SlideImage, form auth.PhoneNumberForm) {
	phoneNumber := strings.TrimSpace(form.PhoneNumber)

	if !phone.IsValidPhoneNumber(phoneNumber) {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.format_err")))
		return
	}

	hasManual, err := slideImage.VerifyManual(form.SlideID)
	if err != nil {
		log.Warn("redis err", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
		return

	}
	if !hasManual {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
		return
	}

	if form.Mode != 2 {
		has, err := models.IsUserByPhoneNumberExist(phoneNumber)
		if err != nil {
			log.Warn("sql err", err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
			return
		}

		if form.Mode == 0 { //注册

			if has {
				ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.already_register")))
				return
			}
		} else { //手机号验证码登录  mode=1   忘记密码 mode=3
			if !has {
				ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.not_register")))
				return
			}

		}

	} else {
		//修改手机号 mode=2 绑定手机
		u, err := models.GetUserByPhoneNumber(phoneNumber)
		if err != nil && !models.IsErrUserNotExist(err) {
			log.Warn("sql err", err)
			ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
			return
		}

		if u != nil {

			if u.ID == ctx.User.ID { //没有修改手机号
				ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.not_modify")))
				return
			} else { //修改的手机已经被别的用户注册
				ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.already_register")))
				return
			}

		}
	}

	redisConn := labelmsg.Get()
	defer redisConn.Close()

	sendTimes, err := phoneService.GetPhoneNumberSendTimes(redisConn, phoneNumber)
	if err != nil && err != redis.ErrNil {
		log.Warn("redis err", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
		return

	}
	if sendTimes >= setting.PhoneService.MaxRetryTimes {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.max_times", strconv.Itoa(setting.PhoneService.MaxRetryTimes))))
		return

	}

	ttl, err := phoneService.GetPhoneCodeTTL(redisConn, phoneNumber)
	if err != nil {
		log.Warn("redis err", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
		return

	}
	if setting.PhoneService.CodeTimeout-ttl < setting.PhoneService.RetryInterval {
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.too_fast")))
		return
	}
	err = phoneService.SendVerifyCode(redisConn, phoneNumber)
	if err != nil {
		log.Warn("send code or redis err", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessage(ctx.Tr("phone.query_err")))
		return
	}

	ctx.JSON(http.StatusOK, models.BaseOKMessage)

}
