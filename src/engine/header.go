package engine

import (
	"net/http"

	"github.com/feed3r/21Updater/src/model"
)

func ExtractEventFromHeader(h *http.Header) string {
	event := model.EventTranslation[h.Get(model.GH_HEADER_EVENT)]

	if event == "" {
		event = h.Get(model.GH_HEADER_EVENT)
	}

	return event
}
