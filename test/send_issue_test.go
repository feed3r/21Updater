package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/feed3r/21Updater/test/test_models"
	"github.com/stretchr/testify/require"
)

func getIssueMessage() string {
	return test_models.ISSUE_HEADER + test_models.ISSUE_BODY
}

func TestSendMessageToServer(t *testing.T) {

	// Create HTTP request
	url := "http://localhost:8080/githubUpdate"
	issueMessage := getIssueMessage()

	// Create request body
	body := bytes.NewBuffer([]byte(issueMessage))

	// Create HTTP request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create request headers
	headers := make(http.Header)

	// Set request headers
	for _, line := range strings.Split(test_models.ISSUE_HEADER, "\n") {
		parts := strings.Split(line, ": ")
		headers.Set(parts[0], parts[1])
	}

	// Set request headers
	req.Header = headers

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print response
	responseBody, err := io.ReadAll(resp.Body)
	require.Nil(t, err)
	require.Equal(t, "Got an error sending message to Telegram Chat: ", string(responseBody))

}
