package engine

import (
	"net/http"
	"strings"

	"github.com/feed3r/21Updater/src/model"
	"github.com/sirupsen/logrus"
)

func ParsePush(h *http.Header, b map[string]interface{}, eventDesc *model.GHEventDescriptor, logger *logrus.Logger) {

	//2 - Action
	eventDesc.Action = "PUSHED"

	//4 - Author
	if pusher, pusherExists := b["pusher"].(map[string]interface{}); pusherExists {
		if name, nameExists := pusher["name"].(string); nameExists {
			eventDesc.Author = name
		}
	}

	//7 - Repository Name
	if repo, repoExists := b["repository"].(map[string]interface{}); repoExists {
		if name, nameExists := repo["name"].(string); nameExists {
			eventDesc.RepoName = name
		}
	}

	if branch, branchExists := b["ref"].(string); branchExists {
		//8 - Branch
		lastSlash := strings.LastIndex(branch, "/")
		if lastSlash >= 0 && lastSlash+1 < len(branch) {
			branch = branch[lastSlash+1:]
		}
		eventDesc.Branch = branch
	}

	//8 - Commits
	if commitArray, commitArrayExists := b["commits"].([]interface{}); commitArrayExists {

		//initialize the commits array
		eventDesc.Commits = make([]model.GHEventCommit, len(commitArray))

		for i, commit := range commitArray {
			commitMap, ok := commit.(map[string]interface{})
			if !ok {
				logger.Errorf("Error parsing commit %d", i)
				continue
			}

			//4 - Author: I need to parse committer first and then extract the username
			if committer, committerExists := commitMap["committer"].(map[string]interface{}); committerExists {
				if username, usernameExists := committer["username"].(string); usernameExists {
					eventDesc.Commits[i].Author = username
				}
			}

			//5 - Message
			if message, messageExists := commitMap["message"].(string); messageExists {
				eventDesc.Commits[i].Message = message
			}

			//6 - Commit URL
			if url, urlExists := commitMap["url"].(string); urlExists {
				eventDesc.Commits[i].URL = url
			}
		}
		if !commitArrayExists || len(eventDesc.Commits) == 0 {
			logger.Info("No commits found in the push event, discarding the event")
			eventDesc.ToBeDiscarded = true
		}
	}

}
