package pipeline

import (
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

const (
	defaultTimeout            = 24 * time.Hour
	groupId             int64 = 1
	groupAccount              = "1"
	organizationId      int64 = 1
	organizationAccount       = "1"
	orgStatus                 = "inUse"
)

const (
	UrlGroup = "/pch"
)

func Redirect(ctx *context.Context) {

	u := ctx.User
	if u == nil {
		log.Error("Failed to get user")
		ctx.JSON(http.StatusInternalServerError, "Failed to get user")
		return
	}

	token, err := GenerateToken(UserClaims{
		UserID:              u.ID,
		UserName:            u.Name,
		GroupID:             groupId,
		GroupAccount:        groupAccount,
		OrganizationId:      organizationId,
		OrganizationAccount: organizationAccount,
		OrgStatus:           orgStatus, // note: 否则, POST 等方法, 可能会出现 ErrTokenFrozen
	}, time.Time{})
	if err != nil {
		log.Error("Failed to generate token: %v", err)
		ctx.JSON(http.StatusInternalServerError, "Failed to generate token")
		return
	}

	mlopsHost := GetHost()
	if len(mlopsHost) == 0 {
		log.Error("mlopsHost is empty")
		ctx.JSON(http.StatusInternalServerError, "mlopsHost is empty")
		return
	}
	target, err := url.Parse(mlopsHost)
	if err != nil {
		log.Error(fmt.Sprintf("url.Parse error: %+v, MlopsHost: '%s'", err, mlopsHost))
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("url.Parse error: %+v, MlopsHost: '%s'", err, mlopsHost))
	}

	realproxy := httputil.NewSingleHostReverseProxy(target)
	realproxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	incomingReq := ctx.Req.Request
	incomingReq.URL.Path = strings.TrimPrefix(incomingReq.URL.Path, UrlGroup)
	incomingReq.Header.Set("Authorization", "Bearer "+token) // set 和 add 的区别

	log.Info("Redirecting to =====> %s", incomingReq.URL)
	log.Info("Redirecting to Host =====> %s", incomingReq.Host)
	realproxy.ServeHTTP(ctx.Resp, incomingReq)
}

type UserClaims struct {
	UserID              int64  `json:"user_id"`
	UserName            string `json:"user_name,omitempty"`
	GroupID             int64  `json:"group_id,omitempty"`
	GroupAccount        string `json:"group_account,omitempty"`
	OrganizationId      int64  `json:"organization_id,omitempty"`
	OrganizationAccount string `json:"organization_account,omitempty"`
	OrgStatus           string `json:"status,omitempty"`
	Role                int    `json:"role,omitempty"`
	Usage               string `json:"usage,omitempty"`
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(UserClaims); ok {
		return jwt.MapClaims{
			UserName:            v.UserName,
			UserId:              v.UserID,
			GroupId:             v.GroupID,
			GroupAccount:        v.GroupAccount,
			OrganizationId:      v.OrganizationId,
			OrganizationAccount: v.OrganizationAccount,
			OrganizationStatus:  v.OrgStatus,
			RoleField:           v.Role,
			"usage":             v.Usage,
		}
	}
	return jwt.MapClaims{}
}

func GenerateToken(uc UserClaims, expireAt time.Time) (token string, err error) {
	now := GetNow()
	if expireAt.IsZero() {
		expireAt = now.Add(defaultTimeout)
	}

	claims := payloadFunc(uc)
	claims["exp"] = expireAt.Unix()
	claims["orig_iat"] = now.Unix()

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(GetJwtSecretKey()))
	if err != nil {
		err = errors.Wrap(err, "SignedString failed")
		return
	}

	return
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetJwtSecretKey() string {
	if setting.MlopsSecret == "" {
		return "jwt secret key"
	}
	return setting.MlopsSecret
}

// supposed to be something like "https://dev-1-37.apulis.com.cn"
func GetHost() string {
	return setting.MlopsHost
}
