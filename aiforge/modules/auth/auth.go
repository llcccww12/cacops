// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"reflect"
	"strings"

	"code.gitea.io/gitea/modules/base"
	"code.gitea.io/gitea/modules/setting"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/auth/sso"
	"code.gitea.io/gitea/modules/validation"

	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
	"gitea.com/macaron/session"
	"github.com/unknwon/com"
)

// IsAPIPath if URL is an api path
func IsAPIPath(url string) bool {
	return strings.HasPrefix(url, "/api/")
}

// SignedInUser returns the user object of signed user.
// It returns a bool value to indicate whether user uses basic auth or not.
func SignedInUser(ctx *macaron.Context, sess session.Store) (*models.User, bool) {
	if !models.HasEngine {
		return nil, false
	}

	checkAutoLogin(ctx, sess)

	// Try to sign in with each of the enabled plugins
	for _, ssoMethod := range sso.Methods() {
		if !ssoMethod.IsEnabled() {
			continue
		}
		user := ssoMethod.VerifyAuthData(ctx, sess)
		if user != nil {
			_, isBasic := ssoMethod.(*sso.Basic)
			return user, isBasic
		}
	}

	return nil, false
}

func checkAutoLogin(ctx *macaron.Context, sess session.Store) {
	uid := sess.Get("uid")
	if uid == nil {
		uname := ctx.GetCookie(setting.CookieUserName)

		u, err := models.GetUserByName(uname)
		if err == nil {

			if val, ok := ctx.GetSuperSecureCookie(
				base.EncodeMD5(u.Rands+u.Passwd), setting.CookieRememberName); ok && val == u.Name {
				sess.Set("uid", u.ID)
			}
		}
	}

}

// AutoLoginner 在 session 恢复后、csrf 中间件之前执行 remember-me 自动登录。
// 这样 csrf 中间件读到的 sess.uid 与后续 Contexter 中使用的一致，
// 避免 session 重建时因 uid 从 0 变为 u.ID 而触发 token 强制轮换导致前端 Invalid csrf token。
func AutoLoginner() macaron.Handler {
	return func(ctx *macaron.Context, sess session.Store) {
		checkAutoLogin(ctx, sess)
	}
}

// Form form binding interface
type Form interface {
	binding.Validator
}

func init() {
	binding.SetNameMapper(com.ToSnakeCase)
}

// AssignForm assign form values back to the template data.
func AssignForm(form interface{}, data map[string]interface{}) {
	typ := reflect.TypeOf(form)
	val := reflect.ValueOf(form)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		fieldName := field.Tag.Get("form")
		// Allow ignored fields in the struct
		if fieldName == "-" {
			continue
		} else if len(fieldName) == 0 {
			fieldName = com.ToSnakeCase(field.Name)
		}

		data[fieldName] = val.Field(i).Interface()
	}
}

func getRuleBody(field reflect.StructField, prefix string) string {
	for _, rule := range strings.Split(field.Tag.Get("binding"), ";") {
		if strings.HasPrefix(rule, prefix) {
			return rule[len(prefix) : len(rule)-1]
		}
	}
	return ""
}

// GetSize get size int form tag
func GetSize(field reflect.StructField) string {
	return getRuleBody(field, "Size(")
}

// GetMinSize get minimal size in form tag
func GetMinSize(field reflect.StructField) string {
	return getRuleBody(field, "MinSize(")
}

// GetMaxSize get max size in form tag
func GetMaxSize(field reflect.StructField) string {
	return getRuleBody(field, "MaxSize(")
}

// GetInclude get include in form tag
func GetInclude(field reflect.StructField) string {
	return getRuleBody(field, "Include(")
}

func validate(errs binding.Errors, data map[string]interface{}, f Form, l macaron.Locale) binding.Errors {
	if errs.Len() == 0 {
		return errs
	}

	data["HasError"] = true
	// If the field with name errs[0].FieldNames[0] is not found in form
	// somehow, some code later on will panic on Data["ErrorMsg"].(string).
	// So initialize it to some default.
	data["ErrorMsg"] = l.Tr("form.unknown_error")
	AssignForm(f, data)

	typ := reflect.TypeOf(f)
	val := reflect.ValueOf(f)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if field, ok := typ.FieldByName(errs[0].FieldNames[0]); ok {
		fieldName := field.Tag.Get("form")
		if fieldName != "-" {
			data["Err_"+field.Name] = true

			trName := field.Tag.Get("locale")
			if len(trName) == 0 {
				trName = l.Tr("form." + field.Name)
			} else {
				trName = l.Tr(trName)
			}

			switch errs[0].Classification {
			case binding.ERR_REQUIRED:
				data["ErrorMsg"] = trName + l.Tr("form.require_error")
			case binding.ERR_ALPHA_DASH:
				data["ErrorMsg"] = trName + l.Tr("form.alpha_dash_error")
			case binding.ERR_ALPHA_DASH_DOT:
				data["ErrorMsg"] = trName + l.Tr("form.alpha_dash_dot_error")
			case validation.ErrGitRefName:
				data["ErrorMsg"] = trName + l.Tr("form.git_ref_name_error")
			case binding.ERR_SIZE:
				data["ErrorMsg"] = trName + l.Tr("form.size_error", GetSize(field))
			case binding.ERR_MIN_SIZE:
				data["ErrorMsg"] = trName + l.Tr("form.min_size_error", GetMinSize(field))
			case binding.ERR_MAX_SIZE:
				data["ErrorMsg"] = trName + l.Tr("form.max_size_error", GetMaxSize(field))
			case binding.ERR_EMAIL:
				data["ErrorMsg"] = trName + l.Tr("form.email_error")
			case binding.ERR_URL:
				data["ErrorMsg"] = trName + l.Tr("form.url_error")
			case binding.ERR_INCLUDE:
				data["ErrorMsg"] = trName + l.Tr("form.include_error", GetInclude(field))
			case validation.ErrGlobPattern:
				data["ErrorMsg"] = trName + l.Tr("form.glob_pattern_error", errs[0].Message)
			case validation.ErrAlphaDashDotChinese:
				data["ErrorMsg"] = trName + l.Tr("form.alpha_dash_dot_chinese_error")
			default:
				data["ErrorMsg"] = l.Tr("form.unknown_error") + " " + errs[0].Classification
			}
			return errs
		}
	}
	return errs
}
