package test

import (
	"fmt"
	"testing"

	"github.com/feed3r/21Updater/src/model"
	"gopkg.in/yaml.v3"
)

func TestYAMLParsing(t *testing.T) {
	// Define the expected YAML data
	yamlData := `
host:
  address: localhost
  port: 8080
  bot_api_key: aaa_bbb_ccc
`

	// Parse the YAML data into a struct
	var parsedData model.Conf
	err := yaml.Unmarshal([]byte(yamlData), &parsedData)
	if err != nil {
		t.Fatalf("Error parsing YAML: %v", err)
	}

	// Verify the parsed data
	expectedConf := model.Conf{
		Host: {
			Address: "localhost",
			Port:    8080,
		},
		BotApiKey: "aaa_bbb_ccc",
	}

	if parsedData["person"] != expectedConf {
		t.Errorf("Parsed data doesn't match the expected result.\nExpected: %+v\nGot: %+v", expectedPerson, parsedData["person"])
	}
}

func main() {
	result := testing.MainStart(nil, nil, main, nil)
	fmt.Println(result.Run())
}
