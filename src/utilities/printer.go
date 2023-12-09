package utilities

import (
	"fmt"
	"net/http"
)

func PrintHeaders(headers http.Header) {
	// Print all the headers
	fmt.Println("HTTP Header:")
	for key, values := range headers {
		fmt.Printf("%s: %s\n", key, values)
	}
}
