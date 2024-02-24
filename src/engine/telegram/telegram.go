package telegram

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/feed3r/21Updater/src/model"
)

func SendTextToTelegramChat(botToken string, chatId string, text string) (*model.TelegramResponse, error) {

	log.Printf("Sending %s to chat_id: %s", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + botToken + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {chatId},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// parse the response
	result := &model.TelegramResponse{}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		log.Printf("could not decode response from telegram: %s", err.Error())
		return nil, err
	}
	return result, nil
}

func ParseTelegramRequest(r *http.Response) (*model.Update, error) {
	var update model.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}
