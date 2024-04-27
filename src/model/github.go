package model

import (
	"bytes"
	"text/template"
)

const PUSH_EVENT = "PUSHED"
const GH_HEADER_EVENT = "X-GitHub-Event"
const TELEGRAM_MESSAGE = `Hey, in our repository {{.RepoName}} a new {{.Event}} has occured, with Action: [{{.Action}}] Author: [{{.Author}}] and Title: [{{.Title}}].
Text says: "{{.Message}}"

You can see the event here: "{{.EventURL}}"`

const TELEGRAM_PUSH_MESSAGE = `Hey, brand new code has been {{.Event}} in our {{.RepoName}} repository and in branch [{{.Branch}}] by [{{.Author}}]
Following the commits info:

{{range .Commits}}

Author: 
"{{.Author}}"
  
Message: 
"{{.Message}}"


Link to the commit details: 
{{.URL}}
------------------------------------------

{{end}}`

var EventConversion map[string]string

func init() {
	EventConversion = map[string]string{
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
	Branch   string
	Event    string
	Action   string
	Author   string
	Title    string
	Message  string
	EventURL string
	Commits  []GHEventCommit
}

const DEFAULT_LANG = "en"

// String returns the string representation of the GHEventDescriptor
func (e *GHEventDescriptor) String(translator *GHEventTranslator) string {

	var tmpl *template.Template
	var err error

	if e.Event == PUSH_EVENT {
		tmpl, err = template.New("telegram_push_message").Parse(translator.PushMessage)
		if err != nil {
			panic(err)
		}
	} else {
		tmpl, err = template.New("telegram_message").Parse(translator.Message)
		if err != nil {
			panic(err)
		}
	}
	var message bytes.Buffer

	err = tmpl.Execute(&message, e)
	if err != nil {
		panic(err)
	}

	return message.String()
}
