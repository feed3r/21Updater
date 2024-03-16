package telegram

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
)

func SendTextToTelegramChat(botToken string, chatId string, text string, logger *logrus.Logger) (*model.TelegramResponse, error) {

	logger.Printf("Sending %s to chat_id: %s", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + botToken + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {chatId},
			"text":    {text},
		})

	if err != nil {
		logger.Error("error when posting text to the chat: " + err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// parse the response
	result := &model.TelegramResponse{}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		logger.Error("could not decode response from telegram: " + err.Error())
		return nil, err
	}
	return result, nil
}

func ParseTelegramRequest(r *http.Response, logger *logrus.Logger) (*model.Update, error) {
	var update model.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		logger.Error("could not decode incoming update " + err.Error())
		return nil, err
	}
	return &update, nil
}
