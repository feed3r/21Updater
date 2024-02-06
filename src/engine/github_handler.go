package engine

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleGithubUpdate(w http.ResponseWriter, r *http.Request) {

	//read the header
	headers := r.Header

	decoder := json.NewDecoder(r.Body)

	var jsonData map[string]interface{}

	err := decoder.Decode(&jsonData)
	if err != nil {
		panic(err)
	}

	log.Println("Received the following JSON: ", jsonData)

	eventDesc := ParseIssue(&headers, jsonData)

	log.Println("Going to send the following message to Telegram chat: ", eventDesc)

	w.Write([]byte(eventDesc.String()))
}
