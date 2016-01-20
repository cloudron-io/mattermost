// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"encoding/json"
	"github.com/mattermost/platform/model"
	"github.com/mattermost/platform/model/oauth"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig("config.json")
}

func TestOAuthProviderLoading(t *testing.T) {
	LoadConfig("config.json")

	tempConfigDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Error creating temporary directory for oauth provider testing: %s", err)
	}
	defer os.RemoveAll(tempConfigDir)

	jsonData, err := json.Marshal(&oauth.OAuthProviderSettings{
		model.SSOSettings{
			Enable:          true,
			Secret:          "mysecret",
			Id:              "myid",
			Scope:           "authorize",
			AuthEndpoint:    "https://mattermost.org/authorize",
			TokenEndpoint:   "https://mattermost.org/oauth/token",
			UserApiEndpoint: "https://mattermost.org/api/v1/userinfo",
		},
		"myusername",
		"myemail",
		"myid",
	})
	if err != nil {
		t.Fatalf("Error marshalling temporary oauth provider json data: %s", err)
	}

	if err := ioutil.WriteFile(filepath.Join(tempConfigDir, "provider1.json"), jsonData, 0600); err != nil {
		t.Fatalf("Error writing temporary oauth provider file: %s", err)
	}

	jsonData, err = json.Marshal(&oauth.OAuthProviderSettings{
		model.SSOSettings{
			Enable:          true,
			Secret:          "mysecret2",
			Id:              "myid2",
			Scope:           "authorize",
			AuthEndpoint:    "https://mattermost.org/authorize",
			TokenEndpoint:   "https://mattermost.org/oauth/token",
			UserApiEndpoint: "https://mattermost.org/api/v1/userinfo",
		},
		"myusername",
		"myemail",
		"myid",
	})
	if err != nil {
		t.Fatalf("Error marshalling second temporary oauth provider json data: %s", err)
	}

	if err := ioutil.WriteFile(filepath.Join(tempConfigDir, "provider2.json"), jsonData, 0600); err != nil {
		t.Fatalf("Error writing second temporary oauth provider file: %s", err)
	}

	jsonData, err = json.Marshal(&oauth.OAuthProviderSettings{
		model.SSOSettings{
			Enable:          false,
			Secret:          "mysecret2",
			Id:              "myid2",
			Scope:           "authorize",
			AuthEndpoint:    "https://mattermost.org/authorize",
			TokenEndpoint:   "https://mattermost.org/oauth/token",
			UserApiEndpoint: "https://mattermost.org/api/v1/userinfo",
		},
		"myusername",
		"myemail",
		"myid",
	})
	if err != nil {
		t.Fatalf("Error marshalling second temporary oauth provider json data: %s", err)
	}

	if err := ioutil.WriteFile(filepath.Join(tempConfigDir, "disabledprovider.json"), jsonData, 0600); err != nil {
		t.Fatalf("Error writing second temporary oauth provider file: %s", err)
	}

	Cfg.OAuthConfigDir = tempConfigDir

	if Cfg.OAuthSettings != nil {
		t.Fatal("Cfg.OAuthSettings unexpectedly set in default model config")
	}

	if err := loadOAuthProviders(Cfg); err != nil {
		t.Fatalf("Error loading test auth providers: %s", err)
	}

	if Cfg.OAuthSettings == nil {
		t.Fatalf("Error loading test oauth providers")
	}

	if _, contained := Cfg.OAuthSettings["disabledprovider"]; contained {
		t.Fatalf("disabledprovider.json should have been ignored")
	}

	if provider1, ok := Cfg.OAuthSettings["provider1"]; ok {
		if provider1.Secret != "mysecret" {
			t.Fatal("provider1 was not loaded correctly")
		}
	} else {
		t.Fatalf("provider1.json was not loaded")
	}

	if provider2, ok := Cfg.OAuthSettings["provider2"]; ok {
		if provider2.Secret != "mysecret2" {
			t.Fatal("provider2 was not loaded correctly")
		}
	} else {
		t.Fatalf("provider2.json was not loaded")
	}
}
