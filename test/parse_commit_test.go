package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/feed3r/21Updater/src/model"
	"github.com/feed3r/21Updater/test/test_models"
	"github.com/feed3r/21Updater/test/utils"
	"github.com/stretchr/testify/require"
)

func TestParseCommit(t *testing.T) {

	headers, err := utils.ReadHeaders(test_models.COMMIT_HEADERS)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var commitJson map[string]interface{}
	json.Unmarshal([]byte(test_models.COMMIT_BODY), &commitJson)

	var eventDesc = new(model.GHEventDescriptor)
	eventDesc.Event = engine.ExtractEventFromHeader(&headers)
	require.Equal(t, "push", strings.ToLower(eventDesc.Event))

	engine.ParseIssue(&headers, commitJson, eventDesc)
	require.Equal(t, test_models.COMMIT_EXPECTED_TEXT, eventDesc.String())

}
