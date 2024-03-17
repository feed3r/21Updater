package main

import (
	"net/http"
	"strconv"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/sirupsen/logrus"
)

func main() {

	var logger = logrus.New()

	//load configuration
	config, err := engine.LoadConfig()
	if err != nil {
		logger.Fatal("Error loading configuration: ", err)
		return
	}

	//load log file
	logger, logFile, err := engine.GetLogger(config)
	if err != nil {
		logger.Fatal("Error loading log file: ", err)
		return
	}
	defer logFile.Close()

	logger.Info("Logger created")

	http.HandleFunc("/githubUpdate", func(w http.ResponseWriter, r *http.Request) {
		engine.HandleGithubUpdate(w, r, config, logger)
	})

	serverAddr := config.Host.Address + ":" + strconv.Itoa(config.Host.Port)
	logger.Info("21Updater server started and listening at " + serverAddr)

	http.ListenAndServe(serverAddr, nil)
}
