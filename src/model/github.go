package model

import (
	"bytes"
	"text/template"
)

const GH_HEADER_EVENT = "X-GitHub-Event"
const TELEGRAM_MESSAGE = `Hey, in our repository {{.RepoName}} a new {{.Event}} has occured, with Action: [{{.Action}}] Author: [{{.Author}}] and Title: [{{.Title}}].
Text says: "{{.Message}}"

You can see the event here: "{{.EventURL}}"`

const TELEGRAM_PUSH_MESSAGE = `"Hey, brand new code has been {{.Event}} in our {{.RepoName}} repository by [{{.Author}}]
Following the commits info:

{{range .Commits}}
- Author: {{.Author}}
  Message: {{.Message}}
  URL: {{.URL}}
{{end}}`

// action translation map
var EventTranslation map[string]string

func init() {
	EventTranslation = map[string]string{
		"issues":       "ISSUE",
		"pull_request": "PULL REQUEST",
		"push":         "PUSHED",
	}
}

type GHEventCommit struct {
	Author  string
	Message string
	URL     string
}

type GHEventDescriptor struct {
	RepoName string
	Event    string
	Action   string
	Author   string
	Title    string
	Message  string
	EventURL string
	Commits  []GHEventCommit
}

// String returns the string representation of the GHEventDescriptor
func (e *GHEventDescriptor) String() string {
	tmpl, err := template.New("telegram_message").Parse(TELEGRAM_MESSAGE)
	if err != nil {
		panic(err)
	}

	var message bytes.Buffer

	err = tmpl.Execute(&message, e)
	if err != nil {
		panic(err)
	}

	return message.String()
}
