package test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/src/engine"
	"github.com/feed3r/21Updater/test/test_models"
	"github.com/stretchr/testify/require"
)

func readHeaders() (http.Header, error) {

	headers := make(http.Header)
	scanner := bufio.NewScanner(bytes.NewBufferString(test_models.ISSUE_HEADER))
	for scanner.Scan() {
		headerLine := scanner.Text()

		if headerLine == "" {
			break
		}

		parts := strings.SplitN(headerLine, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			headers.Add(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headers, nil
}

func rebuildRequest(headers http.Header, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return req, nil
}

func TestRebuildRequest(t *testing.T) {
	headers, err := readHeaders()
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	body, err := os.ReadFile("input_files/test_commit.json")
	if err != nil {
		fmt.Println("Error in reading test file for body: ", err)
	}

	req, err := rebuildRequest(headers, body)
	if err != nil {
		fmt.Println("Error in building the request: ", err)
	}

	fmt.Println("Request rebuilt successfully", req)

}

func TestManageHeaderAction(t *testing.T) {
	headers, err := readHeaders()
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

func TestParseIssue(t *testing.T) {

	headers, err := readHeaders()
	if err != nil {
		fmt.Println("Error in reading test file for headers: ", err)
		return
	}

	var issueJson map[string]interface{}
	json.Unmarshal([]byte(test_models.ISSUE_BODY), &issueJson)

	res := engine.ParseIssue(&headers, issueJson)
	require.Equal(t, test_models.ISSUE_EXPECTED_TEXT, res.String())

}
