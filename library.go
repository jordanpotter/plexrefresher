package plexrefresher

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
)

type LibraryContainer struct {
	XMLName   xml.Name  `xml:"MediaContainer"`
	Libraries []Library `xml:"Directory"`
}

type Library struct {
	XMLName xml.Name `xml:"Directory"`
	Key     string   `xml:"key,attr"`
	Type    string   `xml:"type,attr"`
	Title   string   `xml:"title,attr"`
}

func (pr *PlexRefresher) library(ctx context.Context, title string) (*Library, error) {
	libraries, err := pr.libraries(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get libraries")
	}

	for _, library := range libraries {
		if library.Title == title {
			return &library, nil
		}
	}

	return nil, errors.Errorf("no matching library with title %q", title)
}

func (pr *PlexRefresher) libraries(ctx context.Context) ([]Library, error) {
	url := fmt.Sprintf("%s/library/sections?X-Plex-Token=%s", pr.endpoint, pr.token)
	resp, err := pr.get(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed get request")
	}
	defer resp.Body.Close()

	var container LibraryContainer
	if err = xml.NewDecoder(resp.Body).Decode(&container); err != nil {
		return nil, errors.Wrap(err, "failed to decode HTTP response XML")
	}

	return container.Libraries, nil
}
