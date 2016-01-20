// Copyright (c) 2016 The Qt Company Ltd. All Rights Reserved.
// See License.txt for license information.

package oauth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestOAuthProvider(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Error creating temporary file for oauth provider testing: %s", err)
	}
	defer os.RemoveAll(tempDir)

	testSettings := `{
    "Enable": true,
    "Secret": "mysecret",
    "Id": "myid",
    "Scope": "authorize",
    "AuthEndpoint": "https://mattermost.org/authorize",
    "TokenEndpoint": "https://mattermost.org/oauth/token",
    "UserApiEndpoint": "https://mattermost.org/v1/api/userinfo",
    "DisplayName": "Log into test account",
    "UsernameField": "myusername",
    "EMailField": "myemail",
    "AuthDataField": "myid"
}`
	providerFilePath := filepath.Join(tempDir, "testprovider.json")
	if err := ioutil.WriteFile(providerFilePath, []byte(testSettings), 0644); err != nil {
		t.Fatalf("Error writing testprovider.json file: %s", err)
	}

	providerName, provider, settings, err := LoadOAuthProviderFromSettings(providerFilePath)
	if err != nil {
		t.Fatalf("Unexpected error loading test settings: %s", err)
	}

	if providerName != "testprovider" {
		t.Fatalf("Unexpected provider name %s - expected testprovider", providerName)
	}

	if provider == nil {
		t.Fatal("Unexpected nul provider")
	}

	if settings == nil {
		t.Fatalf("Unexpected nil settings")
	}

	if provider.GetIdentifier() != providerName {
		t.Fatalf("Unexpected provider identifier: %s", provider.GetIdentifier())
	}

	if provider.DisplayName == "" {
		t.Fatalf("Missing display name")
	}

	testUser := `{
    "myusername": "testuser",
    "myemail": "user@host.org",
    "myid": "nickname"
}`

	authData := provider.GetAuthDataFromJson(bytes.NewBufferString(testUser))
	if authData != "nickname" {
		t.Fatalf("Unexpected value for id field %s", authData)
	}

	user := provider.GetUserFromJson(bytes.NewBufferString(testUser))
	if user == nil {
		t.Fatal("Unexpected error parsing test user data")
	}
	if user.Username != "testuser" {
		t.Fatalf("Unexpected user field: %s", user.Username)
	}
	if user.Email != "user@host.org" {
		t.Fatalf("Unexpected email field: %s", user.Email)
	}
	if user.AuthData != "nickname" {
		t.Fatalf("Unexpected id field: %s", user.AuthData)
	}
	if user.AuthService != providerName {
		t.Fatalf("Unexpected auth service field in test user: %s", user.AuthService)
	}
}

func TestInvalidOAuthProvider(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Error creating temporary file for oauth provider testing: %s", err)
	}
	defer os.RemoveAll(tempDir)

	testSettings := `{
    "Enable": true,
    "Secret": "mysecret",
    "Id": "myid",
    "Scope": "authorize",
    "AuthEndpoint": "https://mattermost.org/authorize",
    "TokenEndpoint": "https://mattermost.org/oauth/token",
    "UserApiEndpoint": "https://mattermost.org/v1/api/userinfo"
}`
	providerFilePath := filepath.Join(tempDir, "testprovider.json")
	if err := ioutil.WriteFile(providerFilePath, []byte(testSettings), 0644); err != nil {
		t.Fatalf("Error writing testprovider.json file: %s", err)
	}

	_, _, _, err = LoadOAuthProviderFromSettings(providerFilePath)
	if err == nil {
		t.Fatalf("Unexpected success loading test settings")
	}
	if err.Error() != fmt.Sprintf("Missing UsernameField mapping entry in %s", providerFilePath) {
		t.Fatalf("Unexpected error message when loading incorrect oauth provider settings: %s", err)
	}
}
