package gbdx

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// NewClient returns an oauth2 client configured for token access
func NewClient(ctx context.Context, username string, password string) (*http.Client, error) {

	// oauth2 config with the url
	conf := &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			TokenURL:  BasePath + TokenPath,
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}

	// and get a token
	token, err := conf.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return nil, err
	}
	log.Println("OK: Got GBDX token")

	return conf.Client(ctx, token), nil
}
