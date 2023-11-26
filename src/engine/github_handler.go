package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/feed3r/21Updater/src/model"
)

// parseTelegramRequest handles incoming update from the Telegram web hook
func parseTelegramRequest(r *http.Request) (*model.Update, error) {
	var update model.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
// TODO: Review this
func sendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func writeJSONToFile(filename string, data map[string]interface{}) error {
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

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleGithubUpdate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var json_data map[string]interface{}

	err := decoder.Decode(&json_data)
	if err != nil {
		panic(err)
	}

	//should write the received JSON to file
	err = writeJSONToFile("test_commit.json", json_data)
	if err != nil {
		panic(err)
	}

	fmt.Println(json_data)
	// Parse incoming request
	//TODO
	// var update, err = parseGitHubUpdate(r)
	// if err != nil {
	// 	log.Printf("error parsing update, %s", err.Error())
	// 	return
	// }

	// Sanitize input
	//TODO

	// Build an update message to be sent to the chat
	//TODO

	// Send the message to the Telegram Chat back to Telegram
	// var telegramResponseBody, errTelegram = sendTextToTelegramChat(model.Message.Chat.Id, lyric)
	// if errTelegram != nil {
	// 	log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	// } else {
	// 	log.Printf("punchline %s successfuly distributed to chat id %d", lyric, update.Message.Chat.Id)
	// }
}
