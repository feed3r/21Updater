package test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/feed3r/21Updater/src/model"
	"github.com/feed3r/21Updater/test/test_models"
	"github.com/feed3r/21Updater/test/utils"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

const LANGUAGE_FILEPATH = "./conf/eng.yaml"

func TestParsePush(t *testing.T) {

	logger := logrus.New()
	logger.Out = os.Stderr

	//Load the language file and

	headers, err := utils.ReadHeaders(test_models.PUSH_HEADERS)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var pushJson map[string]interface{}
	json.Unmarshal([]byte(test_models.PUSH_BODY), &pushJson)

	var eventDesc = new(model.GHEventDescriptor)

	eventTranslator, err := utils.ReadDefaultLang()
	if err != nil {
		logger.Fatal("Error in reading the default language file: ", err)
	}

	eventDesc.Event = engine.ExtractEventFromHeader(&headers)
	require.Equal(t, "pushed", strings.ToLower(eventDesc.Event))

	engine.ParsePush(&headers, pushJson, eventDesc, logger)
	require.Equal(t, test_models.PUSH_EXPECTED_TEXT, eventDesc.String(eventTranslator))

}
