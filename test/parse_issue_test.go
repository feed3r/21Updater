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

func TestParseIssue(t *testing.T) {

	logger := logrus.New()
	logger.Out = os.Stderr

	headers, err := utils.ReadHeaders(test_models.ISSUE_HEADER)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var issueJson map[string]interface{}
	json.Unmarshal([]byte(test_models.ISSUE_BODY), &issueJson)

	var eventDesc = new(model.GHEventDescriptor)
	eventDesc.Event = engine.ExtractEventFromHeader(&headers)
	require.Equal(t, "issue", strings.ToLower(eventDesc.Event))

	engine.ParseIssue(&headers, issueJson, eventDesc, logger)
	require.Equal(t, test_models.ISSUE_EXPECTED_TEXT, eventDesc.String())

}
