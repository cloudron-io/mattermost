// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauth

import (
	"encoding/json"
	"github.com/mattermost/platform/einterfaces"
	"github.com/mattermost/platform/model"
	"io"
	// "strings"
)

const (
	USER_AUTH_SERVICE_OAUTH = "oauth"
)

type OAuthProvider struct {
}

type OAuthUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Name     string `json:"name"`
}

func init() {
	provider := &OAuthProvider{}
	einterfaces.RegisterOauthProvider(USER_AUTH_SERVICE_OAUTH, provider)
}

func userFromOAuthUser(glu *OAuthUser) *model.User {
	user := &model.User{}
	username := glu.Username
	user.Username = model.CleanUsername(username)
	// splitName := strings.Split(glu.Name, " ")
	// if len(splitName) == 2 {
	// 	user.FirstName = splitName[0]
	// 	user.LastName = splitName[1]
	// } else if len(splitName) >= 2 {
	// 	user.FirstName = splitName[0]
	// 	user.LastName = strings.Join(splitName[1:], " ")
	// } else {
	// 	user.FirstName = glu.Name
	// }
	user.Email = glu.Email
	user.AuthData = glu.Id
	user.AuthService = USER_AUTH_SERVICE_OAUTH

	return user
}

func OAuthUserFromJson(data io.Reader) *OAuthUser {
	decoder := json.NewDecoder(data)
	var glu OAuthUser
	err := decoder.Decode(&glu)
	if err == nil {
		return &glu
	} else {
		return nil
	}
}

func (glu *OAuthUser) getAuthData() string {
	return glu.Id
}

func (m *OAuthProvider) GetIdentifier() string {
	return USER_AUTH_SERVICE_OAUTH
}

func (m *OAuthProvider) GetUserFromJson(data io.Reader) *model.User {
	return userFromOAuthUser(OAuthUserFromJson(data))
}

func (m *OAuthProvider) GetAuthDataFromJson(data io.Reader) string {
	return OAuthUserFromJson(data).getAuthData()
}
