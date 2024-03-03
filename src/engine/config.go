package engine

import (
	"flag"
	"os"

	"github.com/feed3r/21Updater/src/model"
	"gopkg.in/yaml.v3"
)

// Parse a parameter to provide the config file path
func getConfigFilePath() string {
	confPath := flag.String("conf", "../conf/conf.yaml", "path to the configuration file")
	flag.Parse()
	return *confPath
}

func LoadConfig() (*model.Conf, error) {
	var config model.Conf

	configFilePath := getConfigFilePath()
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
