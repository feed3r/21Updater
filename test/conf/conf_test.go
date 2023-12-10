package conf

import (
	"testing"

	"github.com/feed3r/21Updater/src/model"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestYAMLParsing(t *testing.T) {
	var parsedData model.Conf
	err := yaml.Unmarshal([]byte(BASE_CONF_YAML), &parsedData)
	if err != nil {
		t.Fatalf("Error parsing YAML: %v", err)
	}

	require.Equal(t, parsedData, ExpectedBaseConf)
}
