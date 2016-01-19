// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthqt

import (
	"encoding/json"
	"github.com/mattermost/platform/einterfaces"
	"github.com/mattermost/platform/model"
	"io"
	// "strings"
)

const (
	USER_AUTH_SERVICE_QT = "qt"
)

type QtProvider struct {
}

type QtUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Name     string `json:"name"`
}

func init() {
	provider := &QtProvider{}
	einterfaces.RegisterOauthProvider(USER_AUTH_SERVICE_QT, provider)
}

func userFromQtUser(glu *QtUser) *model.User {
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
	user.AuthService = USER_AUTH_SERVICE_QT

	return user
}

func QtUserFromJson(data io.Reader) *QtUser {
	decoder := json.NewDecoder(data)
	var glu QtUser
	err := decoder.Decode(&glu)
	if err == nil {
		return &glu
	} else {
		return nil
	}
}

func (glu *QtUser) getAuthData() string {
	return glu.Id
}

func (m *QtProvider) GetIdentifier() string {
	return USER_AUTH_SERVICE_QT
}

func (m *QtProvider) GetUserFromJson(data io.Reader) *model.User {
	return userFromQtUser(QtUserFromJson(data))
}

func (m *QtProvider) GetAuthDataFromJson(data io.Reader) string {
	return QtUserFromJson(data).getAuthData()
}
