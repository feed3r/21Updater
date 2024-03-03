package engine

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/feed3r/21Updater/src/engine/telegram"
	"github.com/feed3r/21Updater/src/model"
)

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleGithubUpdate(w http.ResponseWriter, r *http.Request, config *model.Conf) {

	//read the header
	headers := r.Header

	decoder := json.NewDecoder(r.Body)

	var jsonData map[string]interface{}

	err := decoder.Decode(&jsonData)
	if err != nil {
		panic(err)
	}

	log.Println("Received the following JSON: ", jsonData)
	var eventDesc *model.GHEventDescriptor = new(model.GHEventDescriptor)

	eventDesc.Event = ExtractEventFromHeader(&headers)

	switch strings.ToLower(eventDesc.Event) {
	case "issue":
		log.Println("Received an issue event, going to parse it")
		ParseIssue(&headers, jsonData, eventDesc)
	case "pull_request":
		log.Println("Received a pull request event, going to parse it")
		ParsePR(&headers, jsonData, eventDesc)
	default:
		log.Println("Received an event that is not supported: ", eventDesc.Event)
	}

	log.Println("Going to send the following message to Telegram chat: ", eventDesc)

	res, err := telegram.SendTextToTelegramChat(config.BotToken, config.ChatId, eventDesc.String())

	if err != nil {
		w.Write([]byte("Got an error sending message to Telegram Chat: " + err.Error()))
	} else {
		responseByteArray, err := json.Marshal(res)
		if err != nil {
			w.Write([]byte("Got an error marshalling the response: " + err.Error()))
		}

		w.Write(responseByteArray)
	}
}
