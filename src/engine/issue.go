package engine

import (
	"net/http"

	"github.com/feed3r/21Updater/src/model"
)

func ParseIssue(h *http.Header, b map[string]interface{}, eventDesc *model.GHEventDescriptor) {

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

}
