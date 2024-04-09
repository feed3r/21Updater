package engine

import (
	"net/http"

	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
)

func decodePRAction(action string) string {
	switch action {
	case "synchronize":
		return "updated"
	}
	return action
}

func ParsePR(h *http.Header, b map[string]interface{}, eventDesc *model.GHEventDescriptor, logger *logrus.Logger) {

	//2 - Action
	if action, actionExists := b["action"].(string); actionExists {
		eventDesc.Action = decodePRAction(action)
	}

	//3 - Title
	if pr, prExists := b["pull_request"].(map[string]interface{}); prExists {
		if title, titleExists := pr["title"].(string); titleExists {
			eventDesc.Title = title
		}

		//4 - Author
		if user, userExists := pr["user"].(map[string]interface{}); userExists {
			if login, loginExists := user["login"].(string); loginExists {
				eventDesc.Author = login
			}
		}

		//5 - Message
		if body, bodyExists := pr["body"].(string); bodyExists {
			eventDesc.Message = body
		}

		//6 - Event URL
		if url, urlExists := pr["html_url"].(string); urlExists {
			eventDesc.EventURL = url
		}
	}

	//7 - Repository Name
	if repo, repoExists := b["repository"].(map[string]interface{}); repoExists {
		if name, nameExists := repo["name"].(string); nameExists {
			eventDesc.RepoName = name
		}
	}

}
