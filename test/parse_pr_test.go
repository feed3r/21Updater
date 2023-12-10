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

func TestParsePullRequest(t *testing.T) {

	headers, err := utils.ReadHeaders(test_models.PR_HEADER)
	if err != nil {
		fmt.Println("Error in reading test string for headers: ", err)
		return
	}

	var prJson map[string]interface{}
	json.Unmarshal([]byte(test_models.PR_BODY), &prJson)

	res := engine.ParseIssue(&headers, prJson)
	require.Equal(t, test_models.PR_EXPECTED_TEXT, res.String())

}
