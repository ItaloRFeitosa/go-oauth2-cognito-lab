package auth

import (
	"os"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

var (
	OAuth2ClientID     = os.Getenv("OAUTH2_CLIENT_ID")
	OAuth2ClientSecret = os.Getenv("OAUTH2_CLIENT_SECRET")
	OAuth2Scopes       = os.Getenv("OAUTH2_SCOPES")
	OAuth2TokenURL     = os.Getenv("OAUTH2_TOKEN_URL")
)

func ClientCredentialsConfig() *clientcredentials.Config {
	return &clientcredentials.Config{
		ClientID:     OAuth2ClientID,
		ClientSecret: OAuth2ClientSecret,
		Scopes:       SplitBySpace(OAuth2Scopes),
		TokenURL:     OAuth2TokenURL,
	}
}

func SplitBySpace(value string) []string {
	var values []string
	splitted := strings.Split(strings.TrimSpace(value), " ")
	for _, v := range splitted {
		if v != "" {
			values = append(values, v)
		}
	}

	return values
}
