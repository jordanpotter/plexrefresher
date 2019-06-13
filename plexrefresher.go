package plexrefresher

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type PlexRefresher struct {
	httpClient *http.Client
	endpoint   string
	token      string
}

func New(endpoint, token string) (*PlexRefresher, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint cannot be empty")
	} else if _, err := url.Parse(endpoint); err != nil {
		return nil, errors.Wrap(err, "failed to parse endpoint")
	}

	if token == "" {
		return nil, errors.New("takon cannot be empty")
	}

	return &PlexRefresher{
		httpClient: &http.Client{},
		endpoint:   endpoint,
		token:      token,
	}, nil
}
