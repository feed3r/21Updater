package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func WriteJSON(filename string, data map[string]interface{}) error {
	// Codifica la mappa in formato JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Scrivi i dati nel file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func WriteHeaders(filename string, headers http.Header) error {
	// Crea o apri il file in modalità di append
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Scrivi gli header nel file
	fmt.Fprintln(file, "Header del messaggio HTTP:")
	for key, values := range headers {
		for _, value := range values {
			fmt.Fprintf(file, "%s: %s\n", key, value)
		}
	}

	return nil
}
