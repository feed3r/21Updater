package main

import (
	"fmt"
	"net/http"

	"github.com/feed3r/21Updater/src/engine"
)

func main() {
	http.DefaultServeMux.HandleFunc("/githubUpdate", engine.HandleGithubUpdate)

	fmt.Println("21Updater server started and listening at port 9090")
	http.ListenAndServe(":9090", nil)
}
