package plexrefresher

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type PlexRefresher struct {
	endpoint string
	token    string
}

func New(endpoint, token string) (*PlexRefresher, error) {
	pr := PlexRefresher{endpoint, token}

	if err := pr.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to test connection")
	}

	return &pr, nil
}

func (pr *PlexRefresher) Ping() error {
	url := fmt.Sprintf("%s/?X-Plex-Token=%s", pr.endpoint, pr.token)
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed HTTP request")
	}
	defer resp.Body.Close()

	return nil
}

func (pr *PlexRefresher) RefreshLibrary(title string) error {
	library, err := pr.library(title)
	if err != nil {
		return errors.Wrapf(err, "failed to find library %q", title)
	}

	url := fmt.Sprintf("%s/library/sections/%s/refresh?X-Plex-Token=%s", pr.endpoint, library.Key, pr.token)
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "failed HTTP request")
	}
	defer resp.Body.Close()

	return nil
}
