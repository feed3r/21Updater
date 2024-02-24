package main

import (
	"fmt"
	"net/http"

	"github.com/feed3r/21Updater/src/engine"
)

func main() {

	//load configuration
	config, err := engine.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration: ", err)
		return
	}

	http.HandleFunc("/githubUpdate", func(w http.ResponseWriter, r *http.Request) {
		engine.HandleGithubUpdate(w, r, config)
	})

	fmt.Println("21Updater server started and listening at port 9090")
	http.ListenAndServe(":9090", nil)
}
