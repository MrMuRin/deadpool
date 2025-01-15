package auth

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuth struct {
	Config *oauth2.Config
}

func NewGoogleAuth(clientID, clientSecret, redirectURL string) *GoogleAuth {
	return &GoogleAuth{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (g *GoogleAuth) GetUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := g.Config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
