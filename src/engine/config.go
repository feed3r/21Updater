package engine

import (
	"flag"
	"fmt"
	"os"

	"github.com/feed3r/21Updater/src/model"
	"gopkg.in/yaml.v3"
)

var DEFAULT_CONFIG_PATH = "./conf/conf.yaml"

// Parse a parameter to provide the config file path
func getConfigFilePath() string {
	confPath := flag.String("conf", DEFAULT_CONFIG_PATH, "path to the configuration file")
	flag.Parse()
	return *confPath
}

func LoadConfig() (*model.Conf, error) {
	var config model.Conf

	configFilePath := getConfigFilePath()
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("Configuration file not found, please create a valid configuration file as " + DEFAULT_CONFIG_PATH + " or provide a valid path using the -conf flag")
		}
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
