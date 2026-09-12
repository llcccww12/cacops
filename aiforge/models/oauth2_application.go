// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/secret"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/timeutil"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/unknwon/com"
	"golang.org/x/crypto/bcrypt"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type GrantEventType string

const (
	ScopeAdd    GrantEventType = "SCOPE_ADD"
	ScopeRevoke GrantEventType = "SCOPE_REVOKE"

	ApplicationNotVerified  = 1
	ApplicationVerified     = 2
	ApplicationVerifyCancel = 3
)

// OAuth2Application represents an OAuth2 client (RFC 6749)
type OAuth2Application struct {
	ID   int64 `xorm:"pk autoincr"`
	UID  int64 `xorm:"INDEX"`
	User *User `xorm:"-"`

	Name string

	ClientID     string `xorm:"unique"`
	ClientSecret string

	RedirectURIs []string `xorm:"redirect_uris JSON TEXT"`

	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated"`

	AllowedScopes []GrantScope `xorm:"-"`
	VerifyFlag    int          `xorm:"NOT NULL DEFAULT 1"`
}

// TableName sets the table name to `oauth2_application`
func (app *OAuth2Application) TableName() string {
	return "oauth2_application"
}

// PrimaryRedirectURI returns the first redirect uri or an empty string if empty
func (app *OAuth2Application) PrimaryRedirectURI() string {
	if len(app.RedirectURIs) == 0 {
		return ""
	}
	return app.RedirectURIs[0]
}

// LoadUser will load User by UID
func (app *OAuth2Application) LoadUser() (err error) {
	if app.User == nil {
		app.User, err = GetUserByID(app.UID)
	}
	return
}

// ContainsRedirectURI checks if redirectURI is allowed for app
func (app *OAuth2Application) ContainsRedirectURI(redirectURI string) bool {
	return com.IsSliceContainsStr(app.RedirectURIs, redirectURI)
}

// GenerateClientSecret will generate the client secret and returns the plaintext and saves the hash at the database
func (app *OAuth2Application) GenerateClientSecret() (string, error) {
	clientSecret, err := secret.New()
	if err != nil {
		return "", err
	}
	hashedSecret, err := bcrypt.GenerateFromPassword([]byte(clientSecret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	app.ClientSecret = string(hashedSecret)
	if _, err := x.ID(app.ID).Cols("client_secret").Update(app); err != nil {
		return "", err
	}
	return clientSecret, nil
}

// ValidateClientSecret validates the given secret by the hash saved in database
func (app *OAuth2Application) ValidateClientSecret(secret []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(app.ClientSecret), secret) == nil
}

// GetGrantByUserID returns a OAuth2Grant by its user and application ID
func (app *OAuth2Application) GetGrantByUserID(userID int64) (*OAuth2Grant, error) {
	return app.getGrantByUserID(x, userID)
}

func (app *OAuth2Application) getGrantByUserID(e Engine, userID int64) (grant *OAuth2Grant, err error) {
	grant = new(OAuth2Grant)
	if has, err := e.Where("user_id = ? AND application_id = ?", userID, app.ID).Get(grant); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return grant, nil
}

// CreateGrant generates a grant for an user
func (app *OAuth2Application) CreateGrant(userID int64, scopeList ...GrantScope) (*OAuth2Grant, error) {
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return nil, beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
		sess.Close()
	}()

	g, err := app.createGrant(sess, userID, scopeList...)
	if err != nil {
		return nil, err
	}
	err = sess.Commit()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (app *OAuth2Application) createGrant(e Engine, userID int64, scopeList ...GrantScope) (*OAuth2Grant, error) {
	if len(scopeList) == 0 {
		return nil, fmt.Errorf("No scope selected")
	}

	if !app.AreScopesAllowed(scopeList) {
		return nil, fmt.Errorf("Scopes are not allowed")
	}

	grant := &OAuth2Grant{
		ApplicationID: app.ID,
		UserID:        userID,
	}
	_, err := e.Insert(grant)
	if err != nil {
		return nil, err
	}

	_, err = e.Where("application_id = ? and user_id = ?", app.ID, userID).Delete(&OAuth2UserScope{})
	if err != nil {
		return nil, err
	}

	var userGrants = make([]*OAuth2UserScope, 0, len(scopeList))
	var userGrantEvents = make([]*OAuth2GrantEvent, 0, len(scopeList))
	for _, scope := range scopeList {
		userGrants = append(userGrants, &OAuth2UserScope{
			ApplicationID: app.ID,
			UserID:        userID,
			GrantID:       grant.ID,
			Scope:         string(scope),
		})
		userGrantEvents = append(userGrantEvents, &OAuth2GrantEvent{
			ApplicationID:           app.ID,
			UserID:                  userID,
			GrantID:                 grant.ID,
			Scope:                   string(scope),
			EventType:               string(ScopeAdd),
			ApplicationUID:          app.UID,
			ApplicationName:         app.Name,
			ApplicationClientID:     app.ClientID,
			ApplicationRedirectURIs: app.RedirectURIs,
		})
	}
	_, err = e.Insert(userGrants)
	if err != nil {
		return nil, err
	}
	_, err = e.Insert(userGrantEvents)
	if err != nil {
		return nil, err
	}

	return grant, nil
}

// check application permission of scope
func (app *OAuth2Application) AreScopesAllowed(requestScopes []GrantScope) bool {
	allowedScopes := app.GetAllowedScopes()
	if len(allowedScopes) == 0 {
		return false
	}
	set := make(map[GrantScope]struct{}, len(allowedScopes))
	for _, s := range allowedScopes {
		set[s] = struct{}{}
	}
	for _, scope := range requestScopes {
		if _, ok := set[scope]; !ok {
			return false
		}
	}
	return true
}

func (app *OAuth2Application) GetAllowedScopes() []GrantScope {
	if len(app.AllowedScopes) > 0 {
		return app.AllowedScopes
	}
	var scopes = make([]GrantScope, 0)
	err := x.Table("oauth2_application_scope").Select("distinct(scope)").Where("application_id = ?", app.ID).Find(&scopes)
	if err != nil {
		return []GrantScope{}
	}
	if len(scopes) == 0 {
		return DefaultScopes
	}
	return scopes
}

// GetOAuth2ApplicationByClientID returns the oauth2 application with the given client_id. Returns an error if not found.
func GetOAuth2ApplicationByClientID(clientID string) (app *OAuth2Application, err error) {
	return getOAuth2ApplicationByClientID(x, clientID)
}

func getOAuth2ApplicationByClientID(e Engine, clientID string) (app *OAuth2Application, err error) {
	app = new(OAuth2Application)
	has, err := e.Where("client_id = ?", clientID).Get(app)
	if !has {
		return nil, ErrOAuthClientIDInvalid{ClientID: clientID}
	}
	return
}

// GetOAuth2ApplicationByID returns the oauth2 application with the given id. Returns an error if not found.
func GetOAuth2ApplicationByID(id int64) (app *OAuth2Application, err error) {
	return getOAuth2ApplicationByID(x, id)
}

func getOAuth2ApplicationByID(e Engine, id int64) (app *OAuth2Application, err error) {
	app = new(OAuth2Application)
	has, err := e.ID(id).Get(app)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrOAuthApplicationNotFound{ID: id}
	}
	return app, nil
}

// GetOAuth2ApplicationsByUserID returns all oauth2 applications owned by the user
func GetOAuth2ApplicationsByUserID(userID int64) (apps []*OAuth2Application, err error) {
	return getOAuth2ApplicationsByUserID(x, userID)
}

func getOAuth2ApplicationsByUserID(e Engine, userID int64) (apps []*OAuth2Application, err error) {
	apps = make([]*OAuth2Application, 0)
	err = e.Where("uid = ?", userID).Find(&apps)
	return
}

// CreateOAuth2ApplicationOptions holds options to create an oauth2 application
type CreateOAuth2ApplicationOptions struct {
	Name         string
	UserID       int64
	RedirectURIs []string
}

// CreateOAuth2Application inserts a new oauth2 application
func CreateOAuth2Application(opts CreateOAuth2ApplicationOptions) (*OAuth2Application, error) {
	return createOAuth2Application(x, opts)
}

func createOAuth2Application(e Engine, opts CreateOAuth2ApplicationOptions) (*OAuth2Application, error) {
	clientID := uuid.NewV4().String()
	app := &OAuth2Application{
		UID:          opts.UserID,
		Name:         opts.Name,
		ClientID:     clientID,
		RedirectURIs: opts.RedirectURIs,
		VerifyFlag:   ApplicationNotVerified,
	}
	if _, err := e.Insert(app); err != nil {
		return nil, err
	}
	return app, nil
}

// UpdateOAuth2ApplicationOptions holds options to update an oauth2 application
type UpdateOAuth2ApplicationOptions struct {
	ID           int64
	Name         string
	UserID       int64
	RedirectURIs []string
}

// UpdateOAuth2Application updates an oauth2 application
func UpdateOAuth2Application(opts UpdateOAuth2ApplicationOptions) (*OAuth2Application, error) {
	sess := x.NewSession()
	if err := sess.Begin(); err != nil {
		return nil, err
	}
	defer sess.Close()

	app, err := getOAuth2ApplicationByID(sess, opts.ID)
	if err != nil {
		return nil, err
	}
	if app.UID != opts.UserID {
		return nil, fmt.Errorf("UID missmatch")
	}

	app.Name = opts.Name
	app.RedirectURIs = opts.RedirectURIs

	if err = updateOAuth2Application(sess, app); err != nil {
		return nil, err
	}
	app.ClientSecret = ""

	return app, sess.Commit()
}

func updateOAuth2Application(e Engine, app *OAuth2Application) error {
	if _, err := e.ID(app.ID).Update(app); err != nil {
		return err
	}
	return nil
}

func deleteOAuth2Application(sess *xorm.Session, id, userid int64) error {
	if deleted, err := sess.Delete(&OAuth2Application{ID: id, UID: userid}); err != nil {
		return err
	} else if deleted == 0 {
		return fmt.Errorf("cannot find oauth2 application")
	}
	codes := make([]*OAuth2AuthorizationCode, 0)
	// delete correlating auth codes
	if err := sess.Join("INNER", "oauth2_grant",
		"oauth2_authorization_code.grant_id = oauth2_grant.id AND oauth2_grant.application_id = ?", id).Find(&codes); err != nil {
		return err
	}
	codeIDs := make([]int64, 0)
	for _, grant := range codes {
		codeIDs = append(codeIDs, grant.ID)
	}

	if _, err := sess.In("id", codeIDs).Delete(new(OAuth2AuthorizationCode)); err != nil {
		return err
	}

	if _, err := sess.Where("application_id = ?", id).Delete(new(OAuth2Grant)); err != nil {
		return err
	}

	if _, err := sess.Where("application_id = ?", id).Delete(new(OAuth2ApplicationScope)); err != nil {
		return err
	}

	if _, err := sess.Where("application_id = ?", id).Delete(new(OAuth2UserScope)); err != nil {
		return err
	}
	return nil
}

// DeleteOAuth2Application deletes the application with the given id and the grants and auth codes related to it. It checks if the userid was the creator of the app.
func DeleteOAuth2Application(id, userid int64) error {
	sess := x.NewSession()
	if err := sess.Begin(); err != nil {
		return err
	}
	if err := deleteOAuth2Application(sess, id, userid); err != nil {
		return err
	}
	return sess.Commit()
}

// ListOAuth2Applications returns a list of oauth2 applications belongs to given user.
func ListOAuth2Applications(uid int64, listOptions ListOptions) ([]*OAuth2Application, error) {
	sess := x.
		Where("uid=?", uid).
		Desc("id")

	if listOptions.Page != 0 {
		sess = listOptions.setSessionPagination(sess)

		apps := make([]*OAuth2Application, 0, listOptions.PageSize)
		return apps, sess.Find(&apps)
	}

	apps := make([]*OAuth2Application, 0, 5)
	return apps, sess.Find(&apps)
}

//////////////////////////////////////////////////////

// OAuth2AuthorizationCode is a code to obtain an access token in combination with the client secret once. It has a limited lifetime.
type OAuth2AuthorizationCode struct {
	ID                  int64        `xorm:"pk autoincr"`
	Grant               *OAuth2Grant `xorm:"-"`
	GrantID             int64
	Code                string `xorm:"INDEX unique"`
	CodeChallenge       string
	CodeChallengeMethod string
	RedirectURI         string
	ValidUntil          timeutil.TimeStamp `xorm:"index"`
}

// TableName sets the table name to `oauth2_authorization_code`
func (code *OAuth2AuthorizationCode) TableName() string {
	return "oauth2_authorization_code"
}

// GenerateRedirectURI generates a redirect URI for a successful authorization request. State will be used if not empty.
func (code *OAuth2AuthorizationCode) GenerateRedirectURI(state string) (redirect *url.URL, err error) {
	if redirect, err = url.Parse(code.RedirectURI); err != nil {
		return
	}
	q := redirect.Query()
	if state != "" {
		q.Set("state", state)
	}
	q.Set("code", code.Code)
	redirect.RawQuery = q.Encode()
	return
}

// Invalidate deletes the auth code from the database to invalidate this code
func (code *OAuth2AuthorizationCode) Invalidate() error {
	return code.invalidate(x)
}

func (code *OAuth2AuthorizationCode) invalidate(e Engine) error {
	_, err := e.Delete(code)
	return err
}

// ValidateCodeChallenge validates the given verifier against the saved code challenge. This is part of the PKCE implementation.
func (code *OAuth2AuthorizationCode) ValidateCodeChallenge(verifier string) bool {
	return code.validateCodeChallenge(verifier)
}

func (code *OAuth2AuthorizationCode) validateCodeChallenge(verifier string) bool {
	switch code.CodeChallengeMethod {
	case "S256":
		// base64url(SHA256(verifier)) see https://tools.ietf.org/html/rfc7636#section-4.6
		h := sha256.Sum256([]byte(verifier))
		hashedVerifier := base64.RawURLEncoding.EncodeToString(h[:])
		return hashedVerifier == code.CodeChallenge
	case "plain":
		return verifier == code.CodeChallenge
	case "":
		return true
	default:
		// unsupported method -> return false
		return false
	}
}

// GetOAuth2AuthorizationByCode returns an authorization by its code
func GetOAuth2AuthorizationByCode(code string) (*OAuth2AuthorizationCode, error) {
	return getOAuth2AuthorizationByCode(x, code)
}

func getOAuth2AuthorizationByCode(e Engine, code string) (auth *OAuth2AuthorizationCode, err error) {
	auth = new(OAuth2AuthorizationCode)
	if has, err := e.Where("code = ?", code).Get(auth); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	auth.Grant = new(OAuth2Grant)
	if has, err := e.ID(auth.GrantID).Get(auth.Grant); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return auth, nil
}

//////////////////////////////////////////////////////

type GrantScope string

const (
	ScopeUserBase  GrantScope = "user.base"  // 基础信息（用户id，头像等）
	ScopeUserEmail GrantScope = "user.email" // 邮箱地址
	ScopeUserPhone GrantScope = "user.phone" // 手机号
)

var DefaultScopes = []GrantScope{ScopeUserBase}

type OAuth2GrantWithScopes struct {
	ID            int64
	UserID        int64
	Application   *OAuth2Application
	ApplicationID int64
	Counter       int64
	CreatedUnix   timeutil.TimeStamp
	UpdatedUnix   timeutil.TimeStamp
	Scopes        []GrantScope
}

// OAuth2Grant represents the permission of an user for a specifc application to access resources
type OAuth2Grant struct {
	ID            int64              `xorm:"pk autoincr"`
	UserID        int64              `xorm:"INDEX unique(user_application)"`
	Application   *OAuth2Application `xorm:"-"`
	Scopes        []string           `xorm:"-"`
	ApplicationID int64              `xorm:"INDEX unique(user_application)"`
	Counter       int64              `xorm:"NOT NULL DEFAULT 1"`
	CreatedUnix   timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"updated"`
}

// TableName sets the table name to `oauth2_grant`
func (grant *OAuth2Grant) TableName() string {
	return "oauth2_grant"
}

// GenerateNewAuthorizationCode generates a new authorization code for a grant and saves it to the databse
func (grant *OAuth2Grant) GenerateNewAuthorizationCode(redirectURI, codeChallenge, codeChallengeMethod string) (*OAuth2AuthorizationCode, error) {
	return grant.generateNewAuthorizationCode(x, redirectURI, codeChallenge, codeChallengeMethod)
}

func (grant *OAuth2Grant) generateNewAuthorizationCode(e Engine, redirectURI, codeChallenge, codeChallengeMethod string) (code *OAuth2AuthorizationCode, err error) {
	var codeSecret string
	if codeSecret, err = secret.New(); err != nil {
		return &OAuth2AuthorizationCode{}, err
	}
	code = &OAuth2AuthorizationCode{
		Grant:               grant,
		GrantID:             grant.ID,
		RedirectURI:         redirectURI,
		Code:                codeSecret,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: codeChallengeMethod,
	}
	if _, err := e.Insert(code); err != nil {
		return nil, err
	}
	return code, nil
}

// IncreaseCounter increases the counter and updates the grant
func (grant *OAuth2Grant) IncreaseCounter() error {
	return grant.increaseCount(x)
}

func (grant *OAuth2Grant) increaseCount(e Engine) error {
	_, err := e.ID(grant.ID).Incr("counter").Update(new(OAuth2Grant))
	if err != nil {
		return err
	}
	updatedGrant, err := getOAuth2GrantByID(e, grant.ID)
	if err != nil {
		return err
	}
	grant.Counter = updatedGrant.Counter
	return nil
}

func (grant *OAuth2Grant) GetGrantScopes() ([]GrantScope, error) {
	var scopes = make([]GrantScope, 0)
	err := x.Cols("oauth2_user_scope.scope").
		Distinct().
		Table("oauth2_user_scope").
		Join("inner", "oauth2_application_scope", "oauth2_user_scope.application_id = oauth2_application_scope.application_id").
		Where("oauth2_user_scope.grant_id = ? and oauth2_user_scope.scope = oauth2_application_scope.scope", grant.ID).
		Find(&scopes)
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

type OAuth2UserScope struct {
	ID            int64              `xorm:"pk autoincr"`
	GrantID       int64              `xorm:"INDEX"`
	ApplicationID int64              `xorm:"INDEX"`
	UserID        int64              `xorm:"INDEX"`
	CreatedUnix   timeutil.TimeStamp `xorm:"created"`
	Scope         string             `xorm:"NOT NULL"`
}

func (grant *OAuth2UserScope) TableName() string {
	return "oauth2_user_scope"
}

type OAuth2GrantEvent struct {
	ID                      int64 `xorm:"pk autoincr"`
	GrantID                 int64 `xorm:"INDEX"`
	EventType               string
	ApplicationID           int64 `xorm:"INDEX"`
	ApplicationUID          int64 `xorm:""`
	ApplicationName         string
	ApplicationClientID     string             `xorm:""`
	ApplicationRedirectURIs []string           `xorm:"JSON TEXT"`
	UserID                  int64              `xorm:"INDEX"`
	Scope                   string             `xorm:"NOT NULL"`
	CreatedUnix             timeutil.TimeStamp `xorm:"created INDEX"`
}

func (grant *OAuth2GrantEvent) TableName() string {
	return "oauth2_grant_event"
}

type OAuth2ApplicationScope struct {
	ID            int64              `xorm:"pk autoincr"`
	ApplicationID int64              `xorm:"INDEX"`
	Scope         string             `xorm:"NOT NULL"`
	CreatedUnix   timeutil.TimeStamp `xorm:"created"`
	UpdatedUnix   timeutil.TimeStamp `xorm:"updated"`
}

func (grant *OAuth2ApplicationScope) TableName() string {
	return "oauth2_application_scope"
}

// GetOAuth2GrantByID returns the grant with the given ID
func GetOAuth2GrantByID(id int64) (*OAuth2Grant, error) {
	return getOAuth2GrantByID(x, id)
}

func getOAuth2GrantByID(e Engine, id int64) (grant *OAuth2Grant, err error) {
	grant = new(OAuth2Grant)
	if has, err := e.ID(id).Get(grant); err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return
}

// GetOAuth2GrantsByUserID lists all grants of a certain user
func GetOAuth2GrantsByUserID(uid int64) ([]*OAuth2Grant, error) {
	return getOAuth2GrantsByUserID(x, uid)
}

func getOAuth2GrantsByUserID(e Engine, uid int64) ([]*OAuth2Grant, error) {
	type joinedOAuth2Grant struct {
		Grant       *OAuth2Grant       `xorm:"extends"`
		Application *OAuth2Application `xorm:"extends"`
	}
	var results *xorm.Rows
	var err error
	if results, err = e.
		Table("oauth2_grant").
		Where("user_id = ?", uid).
		Join("INNER", "oauth2_application", "application_id = oauth2_application.id").
		Rows(new(joinedOAuth2Grant)); err != nil {
		return nil, err
	}
	defer results.Close()
	grants := make([]*OAuth2Grant, 0)
	for results.Next() {
		joinedGrant := new(joinedOAuth2Grant)
		if err := results.Scan(joinedGrant); err != nil {
			return nil, err
		}
		joinedGrant.Grant.Application = joinedGrant.Application
		grants = append(grants, joinedGrant.Grant)
	}
	return grants, nil
}

// GetOAuth2GrantsByUserID lists all grants of a certain user
func GetOAuth2GrantsWithScopesByUserID(uid int64) ([]*OAuth2Grant, error) {
	return getOAuth2GrantsWithScopesByUserID(x, uid)
}

func getOAuth2GrantsWithScopesByUserID(e Engine, uid int64) ([]*OAuth2Grant, error) {
	type joinedOAuth2Grant struct {
		Grant       *OAuth2Grant       `xorm:"extends"`
		Application *OAuth2Application `xorm:"extends"`
		Scope       *OAuth2UserScope   `xorm:"extends"`
	}
	var results *xorm.Rows
	var err error
	if results, err = e.
		Table("oauth2_grant").
		Where("oauth2_grant.user_id = ?", uid).
		Join("INNER", "oauth2_application", "application_id = oauth2_application.id").
		Join("INNER", "oauth2_user_scope", "oauth2_grant.id = oauth2_user_scope.grant_id").
		Rows(new(joinedOAuth2Grant)); err != nil {
		return nil, err
	}
	defer results.Close()
	grants := make([]*OAuth2Grant, 0)
	grantScopeMap := make(map[int64]*OAuth2Grant, 0)
	for results.Next() {
		joinedGrant := new(joinedOAuth2Grant)
		if err := results.Scan(joinedGrant); err != nil {
			return nil, err
		}
		if _, exsits := grantScopeMap[joinedGrant.Grant.ID]; exsits {
			if len(grantScopeMap[joinedGrant.Grant.ID].Scopes) == 0 {
				grantScopeMap[joinedGrant.Grant.ID].Scopes = []string{joinedGrant.Scope.Scope}
			} else {
				grantScopeMap[joinedGrant.Grant.ID].Scopes = append(grantScopeMap[joinedGrant.Grant.ID].Scopes, joinedGrant.Scope.Scope)
			}
			continue
		}

		joinedGrant.Grant.Application = joinedGrant.Application
		joinedGrant.Grant.Scopes = []string{joinedGrant.Scope.Scope}
		grants = append(grants, joinedGrant.Grant)
		grantScopeMap[joinedGrant.Grant.ID] = joinedGrant.Grant
	}
	return grants, nil
}

// RevokeOAuth2Grant deletes the grant with grantID and userID
func RevokeOAuth2Grant(grantID, userID int64) error {
	return revokeOAuth2Grant(x, grantID, userID)
}

func revokeOAuth2Grant(e Engine, grantID, userID int64) error {
	//判断授权存不存在
	grant := &OAuth2Grant{}
	exists, err := e.ID(grantID).Get(grant)
	if !exists || grant.ID == 0 {
		return nil
	}
	if err != nil {
		return err
	}

	//判断应用是否存在
	app := &OAuth2Application{}
	isAppExists, err := e.ID(grant.ApplicationID).Get(app)
	if !isAppExists || app.ID == 0 {
		return nil
	}
	if err != nil {
		return err
	}

	_, err = e.Delete(&OAuth2Grant{ID: grantID, UserID: userID})
	if err != nil {
		return err
	}

	scopeList := make([]OAuth2UserScope, 0)
	err = e.Where("grant_id = ?", grantID).Find(&scopeList)
	if err != nil {
		//查询失败不进行后续操作也不影响功能，所以直接返回
		log.Error("revokeOAuth2Grant find scopes err.grantID = %d err= %v", grantID, err)
		return nil
	}

	if len(scopeList) > 0 {
		e.Delete(&OAuth2UserScope{GrantID: grantID})
		var userGrantEvents = make([]*OAuth2GrantEvent, 0, len(scopeList))
		for _, scope := range scopeList {
			userGrantEvents = append(userGrantEvents, &OAuth2GrantEvent{
				ApplicationID:           grant.ApplicationID,
				UserID:                  userID,
				GrantID:                 grant.ID,
				Scope:                   scope.Scope,
				EventType:               string(ScopeRevoke),
				ApplicationUID:          app.UID,
				ApplicationName:         app.Name,
				ApplicationClientID:     app.ClientID,
				ApplicationRedirectURIs: app.RedirectURIs,
			})
		}
		_, err = e.Insert(userGrantEvents)
		if err != nil {
			log.Error("revokeOAuth2Grant insert grant events) err.grantID = %d err= %v", grantID, err)
		}

	}

	return nil
}

//////////////////////////////////////////////////////////////

// OAuth2TokenType represents the type of token for an oauth application
type OAuth2TokenType int

const (
	// TypeAccessToken is a token with short lifetime to access the api
	TypeAccessToken OAuth2TokenType = 0
	// TypeRefreshToken is token with long lifetime to refresh access tokens obtained by the client
	TypeRefreshToken = iota
)

// OAuth2Token represents a JWT token used to authenticate a client
type OAuth2Token struct {
	GrantID int64           `json:"gnt"`
	Type    OAuth2TokenType `json:"tt"`
	Counter int64           `json:"cnt,omitempty"`
	jwt.StandardClaims
}

// ParseOAuth2Token parses a singed jwt string
func ParseOAuth2Token(jwtToken string) (*OAuth2Token, error) {
	parsedToken, err := jwt.ParseWithClaims(jwtToken, &OAuth2Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing algo: %v", token.Header["alg"])
		}
		return setting.OAuth2.JWTSecretBytes, nil
	})
	if err != nil {
		return nil, err
	}
	var token *OAuth2Token
	var ok bool
	if token, ok = parsedToken.Claims.(*OAuth2Token); !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

// SignToken signs the token with the JWT secret
func (token *OAuth2Token) SignToken() (string, error) {
	token.IssuedAt = time.Now().Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, token)
	return jwtToken.SignedString(setting.OAuth2.JWTSecretBytes)
}

type OAuthContext struct {
	UserID        int64
	ApplicationID int64
	GrantID       int64
	Scopes        []GrantScope
}

type SearchOauth2ApplicationReq struct {
	ListOptions
	Keyword    string
	VerifyFlag int
	ScopeList  []GrantScope
}

type EditOauth2ApplicationReq struct {
	ID         int64
	VerifyFlag int
	ScopeList  string
}

type Oauth2Application4Show struct {
	ID            int64
	UserID        int64
	UserName      string
	ClientID      string
	Name          string
	VerifyFlag    int
	AllowedScopes []GrantScope
	CreatedUnix   timeutil.TimeStamp
	UpdatedUnix   timeutil.TimeStamp
}

func (Oauth2Application4Show) TableName() string {
	return "oauth2_application"
}

func SearchOauth2ApplicationList(opts SearchOauth2ApplicationReq) ([]*Oauth2Application4Show, int64, error) {
	if opts.Page < 1 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 || opts.PageSize > 100 {
		opts.PageSize = 20
	}

	query := builder.NewCond()
	if opts.Keyword != "" {
		query = query.And(builder.Or(builder.Like{"oauth2_application.name", "%" + opts.Keyword + "%"}, builder.Eq{"oauth2_application.client_id": opts.Keyword}))
	}
	if opts.VerifyFlag > 0 {
		query = query.And(builder.Eq{"oauth2_application.verify_flag": opts.VerifyFlag})
	}

	//默认scope不用当筛选条件，因为很多历史数据没有这个字段，默认所有应用都有
	opts.ScopeList = filterDefaltScopes(opts.ScopeList)

	if len(opts.ScopeList) > 0 {
		distinctScopes := make(map[GrantScope]bool)
		for _, scope := range opts.ScopeList {
			distinctScopes[scope] = true
		}

		query = query.And(builder.In("oauth2_application.id",
			builder.Select("application_id").
				From("oauth2_application_scope").
				Where(builder.In("scope", opts.ScopeList)). // 先过滤相关scope
				GroupBy("application_id").
				Having(fmt.Sprintf("COUNT(DISTINCT scope) = %d", len(distinctScopes)))))
	}

	totalCount, err := x.Where(query).Count(new(Oauth2Application4Show))
	if err != nil {
		return nil, 0, err
	}

	apps := make([]*Oauth2Application4Show, 0, opts.PageSize)
	var cols = "oauth2_application.id as id,oauth2_application.client_id,oauth2_application.name," +
		"oauth2_application.verify_flag,oauth2_application.updated_unix,oauth2_application.created_unix," +
		"public.user.id as user_id,public.user.name as user_name"
	err = x.Cols(cols).Join("inner", "public.user", "oauth2_application.uid = public.user.id").Where(query).Limit(opts.PageSize, opts.PageSize*(opts.Page-1)).OrderBy("updated_unix desc").Find(&apps)
	if err != nil {
		return nil, 0, err
	}
	if apps == nil {
		return nil, 0, nil
	}

	//获取allowed_scopes
	appIds := make([]int64, 0, len(apps))
	for _, app := range apps {
		appIds = append(appIds, app.ID)
	}
	var appScopes = make([]OAuth2ApplicationScope, 0)
	if len(appIds) > 0 {
		err = x.Where(builder.In("application_id", appIds)).Find(&appScopes)
		if err != nil {
			return nil, 0, err
		}
	}

	for i := 0; i < len(apps); i++ {
		var scopes = make([]GrantScope, 0)
		for _, appScope := range appScopes {
			if apps[i].ID == appScope.ApplicationID {
				scopes = append(scopes, GrantScope(appScope.Scope))
			}
		}
		if len(scopes) == 0 {
			scopes = DefaultScopes
		}
		apps[i].AllowedScopes = scopes
	}

	return apps, totalCount, nil
}

func filterDefaltScopes(sliceB []GrantScope) []GrantScope {

	set := make(map[GrantScope]bool)
	r := make([]GrantScope, 0)
	for _, v := range DefaultScopes {
		set[v] = true
	}

	for _, v := range sliceB {
		if !set[v] {
			r = append(r, v)
		}
	}
	return r
}

func UpdateApplicationVerifyStatusAndScopes(app OAuth2Application) error {
	if app.ID <= 0 {
		return fmt.Errorf("param error")
	}
	var err error
	sess := x.NewSession()
	if beginErr := sess.Begin(); beginErr != nil {
		return beginErr
	}

	defer func() {
		if err != nil {
			sess.Rollback()
		}
	}()

	if app.VerifyFlag > 0 {
		_, err = sess.ID(app.ID).Cols("verify_flag").Update(&app)
		if err != nil {
			return err
		}
	}

	_, err = sess.Where("application_id = ?", app.ID).Delete(&OAuth2ApplicationScope{})
	if err != nil {
		return err
	}
	var appScopes = make([]OAuth2ApplicationScope, 0)
	for _, scope := range app.AllowedScopes {
		appScopes = append(appScopes, OAuth2ApplicationScope{
			ApplicationID: app.ID,
			Scope:         string(scope),
		})
	}
	_, err = sess.Insert(&appScopes)
	if err != nil {
		return err
	}
	return sess.Commit()
}
