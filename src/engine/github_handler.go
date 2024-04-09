package engine

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/feed3r/21Updater/src/engine/telegram"
	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
)

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleGithubUpdate(w http.ResponseWriter, r *http.Request, config *model.Conf, logger *logrus.Logger) {

	//read the header
	headers := r.Header

	decoder := json.NewDecoder(r.Body)

	var jsonData map[string]interface{}

	err := decoder.Decode(&jsonData)
	if err != nil {
		panic(err)
	}

	logger.Info("Received the following JSON: ", jsonData)
	var eventDesc *model.GHEventDescriptor = new(model.GHEventDescriptor)

	eventDesc.Event = ExtractEventFromHeader(&headers)

	switch strings.ToLower(eventDesc.Event) {
	case "issue":
		logger.Info("Received an issue event, going to parse it")
		ParseIssue(&headers, jsonData, eventDesc, logger)
	case "pull_request":
		logger.Info("Received a pull request event, going to parse it")
		ParsePR(&headers, jsonData, eventDesc, logger)
	case "push":
		logger.Info("Received a push event, going to parse it")
		ParsePush(&headers, jsonData, eventDesc, logger)
	default:
		logger.Warn("Received an event that is not supported: ", eventDesc.Event)
	}

	logger.Info("Going to send the following message to Telegram chat: ", eventDesc)

	res, err := telegram.SendTextToTelegramChat(config.BotToken, config.ChatId, eventDesc.String(), logger)

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
