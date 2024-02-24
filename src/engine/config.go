package engine

import (
	"os"

	"github.com/feed3r/21Updater/src/model"
	"gopkg.in/yaml.v3"
)

func LoadConfig() (*model.Conf, error) {
	var config model.Conf

	data, err := os.ReadFile("../conf/conf.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
