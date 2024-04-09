package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/feed3r/21Updater/test/test_models"
	"github.com/feed3r/21Updater/test/utils"
	"github.com/stretchr/testify/require"
)

func TestRebuildRequest(t *testing.T) {
	headers, err := utils.ReadHeaders(test_models.PUSH_HEADERS)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	req, err := utils.RebuildRequest(headers, []byte(test_models.PUSH_BODY))
	if err != nil {
		fmt.Println("Error in building the request: ", err)
	}

	fmt.Println("Request rebuilt successfully", req)

}

func TestManageHeaderAction(t *testing.T) {
	headers, err := utils.ReadHeaders(test_models.ISSUE_HEADER)
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	header_event := headers.Get("X-GitHub-Event")
	require.Equal(t, "issues", header_event)
}

func TestReadJson(t *testing.T) {

	var commitJson map[string]interface{}
	json.Unmarshal([]byte(test_models.ISSUE_BODY), &commitJson)

	login := ""

	if sender, senderExists := commitJson["sender"].(map[string]interface{}); senderExists {
		login = sender["login"].(string)
	}

	require.Equal(t, "feed3r", login)

}
