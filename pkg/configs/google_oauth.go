package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

func GetGoogleAuthConfig() *oauth2.Config {
	googleConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:3000/api/v1/authenticate/google/redirect",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return googleConfig
}
