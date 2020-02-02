package provider

import (
	user "github.com/jbltx/rlauncher/user/model"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleProvider ...
type GoogleProvider struct {
	userInfoURL string
	oauth       *oauth2.Config
}

// NewGoogleProvider ...
func NewGoogleProvider() *GoogleProvider {
	return &GoogleProvider{
		userInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
		oauth: &oauth2.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

// OAuth2Config ...
func (p *GoogleProvider) OAuth2Config() *oauth2.Config {
	return p.oauth
}

// UserInfoURL ...
func (p *GoogleProvider) UserInfoURL() string {
	return p.userInfoURL
}

// ParseUserInfo ...
func (p *GoogleProvider) ParseUserInfo(content []byte) (*user.User, error) {
	// TODO
	return nil, nil
}
