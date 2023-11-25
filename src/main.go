package main

import (
	"net/http"

	"github.com/feed3r/21Updater/src/engine"
)

func main() {
	http.DefaultServeMux.HandleFunc("/githubUpdate", engine.HandleGithubUpdate)

	http.ListenAndServe(":8080", nil)
}
