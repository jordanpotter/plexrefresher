package plexrefresher

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (pr *PlexRefresher) Refresh(ctx context.Context, libraryTitle string) error {
	library, err := pr.library(ctx, libraryTitle)
	if err != nil {
		return errors.Wrapf(err, "failed to find library %q", libraryTitle)
	}

	url := fmt.Sprintf("%s/library/sections/%s/refresh?X-Plex-Token=%s", pr.endpoint, library.Key, pr.token)
	resp, err := pr.get(ctx, url)
	if err != nil {
		return errors.Wrap(err, "failed get request")
	}
	defer resp.Body.Close()

	return nil
}
