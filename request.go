package plexrefresher

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

func (pr *PlexRefresher) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req = req.WithContext(ctx)

	resp, err := pr.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform request")
	}

	return resp, nil
}
