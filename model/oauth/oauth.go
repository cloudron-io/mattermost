// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/mattermost/platform/model"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	USER_AUTH_SERVICE_OAUTH = "oauth"
)

type OAuthProvider struct {
	id            string
	usernameField string
	emailField    string
	authDataField string
}

type OAuthProviderSettings struct {
	model.SSOSettings
	UsernameField string
	EMailField    string
	AuthDataField string
}

func userDataFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)
	userData := make(map[string]string)
	err := decoder.Decode(&userData)
	if err == nil {
		return userData
	} else {
		return nil
	}
}

func (m *OAuthProvider) GetIdentifier() string {
	return m.id
}

func (m *OAuthProvider) GetUserFromJson(data io.Reader) *model.User {
	userData := userDataFromJson(data)
	if userData == nil {
		return nil
	}

	user := &model.User{}
	username := userData[m.usernameField]
	user.Username = model.CleanUsername(username)
	user.Email = userData[m.emailField]
	user.AuthData = userData[m.authDataField]
	user.AuthService = m.id
	return user
}

func (m *OAuthProvider) GetAuthDataFromJson(data io.Reader) string {
	userData := userDataFromJson(data)
	if userData == nil {
		return ""
	}
	return userData[m.authDataField]
}

func LoadOAuthProviderFromSettings(settingsJSONFile string) (providerName string, provider *OAuthProvider, settings *model.SSOSettings, err error) {
	var contents []byte
	contents, err = ioutil.ReadFile(settingsJSONFile)
	if err != nil {
		return
	}
	oauthSettings := &OAuthProviderSettings{}
	if err = json.Unmarshal(contents, &oauthSettings); err != nil {
		err = fmt.Errorf("Error reading oauth provider file %s: %s", settingsJSONFile, err)
		return
	}

	ssoSettings := oauthSettings.SSOSettings
	settings = &ssoSettings
	fileName := filepath.Base(settingsJSONFile)
	providerName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	if oauthSettings.UsernameField == "" {
		err = fmt.Errorf("Missing UsernameField mapping entry in %s", settingsJSONFile)
		return
	}

	provider = &OAuthProvider{
		id:            providerName,
		usernameField: oauthSettings.UsernameField,
		emailField:    oauthSettings.EMailField,
		authDataField: oauthSettings.AuthDataField,
	}

	return
}
