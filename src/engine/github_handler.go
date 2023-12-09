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

func writeHeadersToFile(filename string, headers http.Header) error {
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

func ExtractEventFromHeader(h *http.Header) string {
	event := model.EventTranslation[h.Get(model.GH_HEADER_EVENT)]

	if event == "" {
		event = h.Get(model.GH_HEADER_EVENT)
	}

	return event
}

func ParseIssue(h *http.Header, b map[string]interface{}) *model.GHEventDescriptor {

	eventDesc := model.GHEventDescriptor{}

	//1 - Event
	eventDesc.Event = ExtractEventFromHeader(h)

	//2 - Action
	if action, actionExists := b["action"].(string); actionExists {
		eventDesc.Action = action
	}

	//3 - Title
	if issue, issueExists := b["issue"].(map[string]interface{}); issueExists {
		if title, titleExists := issue["title"].(string); titleExists {
			eventDesc.Title = title
		}

		//4 - Author
		if user, userExists := issue["user"].(map[string]interface{}); userExists {
			if login, loginExists := user["login"].(string); loginExists {
				eventDesc.Author = login
			}
		}

		//5 - Message
		if body, bodyExists := issue["body"].(string); bodyExists {
			eventDesc.Message = body
		}
	}

	//6 - Repository Name
	if repo, repoExists := b["repository"].(map[string]interface{}); repoExists {
		if name, nameExists := repo["name"].(string); nameExists {
			eventDesc.RepoName = name
		}
	}

	return &eventDesc
}

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleGithubUpdate(w http.ResponseWriter, r *http.Request) {

	//read the header
	headers := r.Header

	// Print all the headers
	fmt.Println("HTTP Header:")
	for key, values := range headers {
		fmt.Printf("%s: %s\n", key, values)
	}

	decoder := json.NewDecoder(r.Body)

	var jsonData map[string]interface{}

	err := decoder.Decode(&jsonData)
	if err != nil {
		panic(err)
	}

	eventDesc := ParseIssue(&headers, jsonData)

	fmt.Println("Going to send the following message to Telegram chat: ", eventDesc)

	w.Write([]byte(eventDesc.String()))
}
