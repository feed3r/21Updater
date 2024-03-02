package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	serverAddr := config.Host.Address + ":" + strconv.Itoa(config.Host.Port)
	fmt.Println("21Updater server started and listening at " + serverAddr)

	http.ListenAndServe(serverAddr, nil)
}
