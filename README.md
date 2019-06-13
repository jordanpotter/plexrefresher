# Plex Refresher
Simple tool to trigger a Plex library refresh.

## Command Line
```
go run cmd/plexrefresher/main.go -endpoint "http://127.0.0.1:32400" -token "<your_plex_token>" -library "Movies"
```

## Library
```go
pr, err := plexrefresher.New(endpoint, token)
if err != nil {
	log.Fatalf("Unexpected error while creating Plex refresher: %v", err)
}

err = pr.Refresh(context.Background(), library)
if err != nil {
	log.Fatalf("Unexpected error while refreshing Plex library %q: %v", library, err)
}
```
