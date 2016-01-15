// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthcloudron

import (
	"encoding/json"
	"github.com/mattermost/platform/einterfaces"
	"github.com/mattermost/platform/model"
	"io"
	// "strings"
)

const (
	USER_AUTH_SERVICE_CLOUDRON = "cloudron"
)

type CloudronProvider struct {
}

type CloudronUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Name     string `json:"name"`
}

func init() {
	provider := &CloudronProvider{}
	einterfaces.RegisterOauthProvider(USER_AUTH_SERVICE_CLOUDRON, provider)
}

func userFromCloudronUser(glu *CloudronUser) *model.User {
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
	user.AuthService = USER_AUTH_SERVICE_CLOUDRON

	return user
}

func CloudronUserFromJson(data io.Reader) *CloudronUser {
	decoder := json.NewDecoder(data)
	var glu CloudronUser
	err := decoder.Decode(&glu)
	if err == nil {
		return &glu
	} else {
		return nil
	}
}

func (glu *CloudronUser) getAuthData() string {
	return glu.Id
}

func (m *CloudronProvider) GetIdentifier() string {
	return USER_AUTH_SERVICE_CLOUDRON
}

func (m *CloudronProvider) GetUserFromJson(data io.Reader) *model.User {
	return userFromCloudronUser(CloudronUserFromJson(data))
}

func (m *CloudronProvider) GetAuthDataFromJson(data io.Reader) string {
	return CloudronUserFromJson(data).getAuthData()
}
