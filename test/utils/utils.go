package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/feed3r/21Updater/src/model"
	"gopkg.in/yaml.v3"
)

const DEFAULT_LANG_FILE = "../lang/en.yaml"

func ReadHeaders(headers_text string) (http.Header, error) {

	headers := make(http.Header)
	scanner := bufio.NewScanner(bytes.NewBufferString(headers_text))
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

func RebuildRequest(headers http.Header, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	return req, nil
}

func ReadDefaultLang() (*model.GHEventTranslator, error) {
	langFilePath := DEFAULT_LANG_FILE
	var eventTranslator model.GHEventTranslator

	data, err := os.ReadFile(langFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("Cannot find the requested language file " + langFilePath + ", please check that the file exists")
		}
		return nil, err
	}

	err = yaml.Unmarshal([]byte(data), &eventTranslator)
	if err != nil {
		return nil, err
	}

	return &eventTranslator, nil
}
