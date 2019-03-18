package api

import (
	"net/http"

	"github.com/afirth/fm-test/gbdx"
	"golang.org/x/oauth2"
)

// Client must be initialized for handlers which call external services. These handlers are methods of the client. A client is safe for concurrent reuse across goroutines
type Client struct {
	gbdx *http.Client
}

// NewClient returns an initialized client
func NewClient(username string, password string) (c *Client, err error) {
	c = new(Client)
	c.gbdx, err = gbdx.NewClient(oauth2.NoContext, username, password)
	if err != nil {
		return nil, err
	}
	return
}
