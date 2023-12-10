package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/feed3r/21Updater/test/test_models"
	"github.com/feed3r/21Updater/test/utils"
	"github.com/stretchr/testify/require"
)

func TestParseIssue(t *testing.T) {

	headers, err := utils.ReadHeaders(test_models.ISSUE_HEADER)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var issueJson map[string]interface{}
	json.Unmarshal([]byte(test_models.ISSUE_BODY), &issueJson)

	res := engine.ParseIssue(&headers, issueJson)
	require.Equal(t, test_models.ISSUE_EXPECTED_TEXT, res.String())

}
