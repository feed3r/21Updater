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

func TestParsePush(t *testing.T) {

	logger := logrus.New()
	logger.Out = os.Stderr

	headers, err := utils.ReadHeaders(test_models.PUSH_HEADERS)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var pushJson map[string]interface{}
	json.Unmarshal([]byte(test_models.PUSH_BODY), &pushJson)

	var eventDesc = new(model.GHEventDescriptor)
	eventDesc.Event = engine.ExtractEventFromHeader(&headers)
	require.Equal(t, "pushed", strings.ToLower(eventDesc.Event))

	engine.ParsePush(&headers, pushJson, eventDesc, logger)
	require.Equal(t, test_models.PUSH_EXPECTED_TEXT, eventDesc.String())

}

func TestParseFakePush(t *testing.T) {

	logger := logrus.New()
	logger.Out = os.Stderr

	headers, err := utils.ReadHeaders(test_models.FAKE_PUSH_HEADER)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var pushJson map[string]interface{}
	json.Unmarshal([]byte(test_models.FAKE_PUSH_BODY), &pushJson)

	var eventDesc = new(model.GHEventDescriptor)
	eventDesc.Event = "push"
	require.Equal(t, "push", strings.ToLower(eventDesc.Event))

	engine.ParsePush(&headers, pushJson, eventDesc, logger)
	require.Equal(t, true, eventDesc.ToBeDiscarded)

}
