package engine

import (
	"flag"
	"fmt"
	"os"

	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var DEFAULT_CONFIG_PATH = "./conf/conf.yaml"
var DEFAULT_LANG_PATH = "./lang/"

// Parse a parameter to provide the config file path
func getConfigFilePath() string {
	confPath := flag.String("conf", DEFAULT_CONFIG_PATH, "path to the configuration file")
	flag.Parse()
	return *confPath
}

func GetLogger(conf *model.Conf) (*logrus.Logger, *os.File, error) {

	log := logrus.New()

	// If conf.Logfile is empty, the logger will write to the standard output
	if conf.Logfile == "" {
		return log, nil, nil
	}

	file, err := os.OpenFile(conf.Logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return log, nil, err
	}

	// Imposta il file come output del logger
	log.SetOutput(file)

	return log, file, nil
}

func readConf(configFilePath string) (*model.Conf, error) {
	var config model.Conf

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

func readLangTranslator(config *model.Conf) error {
	langFilePath := DEFAULT_LANG_PATH + config.Language + ".yaml"
	data, err := os.ReadFile(langFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("Cannot find the requested language file " + langFilePath + ", please check that the file exists")
		}
		return err
	}

	err = yaml.Unmarshal([]byte(data), &config.Translator)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (*model.Conf, error) {

	configFilePath := getConfigFilePath()
	config, err := readConf(configFilePath)
	if err != nil {
		return nil, err
	}

	err = readLangTranslator(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
