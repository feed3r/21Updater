package utils

import (
	"bufio"
	"bytes"
	"net/http"
	"strings"
)

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
