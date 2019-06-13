package main

import (
	"context"
	"flag"
	"log"

	"github.com/jordanpotter/plexrefresher"
)

var (
	endpoint string
	token    string
	library  string
)

func init() {
	flag.StringVar(&endpoint, "endpoint", "", "API endpoint")
	flag.StringVar(&token, "token", "", "Authentication token")
	flag.StringVar(&library, "library", "", "Library to refresh")
	flag.Parse()
}

func main() {
	pr, err := plexrefresher.New(endpoint, token)
	if err != nil {
		log.Fatalf("Unexpected error while creating Plex refresher: %v", err)
	}

	err = pr.Refresh(context.Background(), library)
	if err != nil {
		log.Fatalf("Unexpected error while refreshing Plex library %q: %v", library, err)
	}
}
