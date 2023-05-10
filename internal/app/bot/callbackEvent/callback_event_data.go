package callbackEvent

import (
	"encoding/json"
	"errors"
)

var allowedTypes = map[string]struct{}{
	"show_snackbar": {},
	"open_link":     {},
	"open_app":      {},
}

type CallbackEvent struct {
	Type    string `json:"type"`
	Text    string `json:"text,omitempty"`
	Link    string `json:"link,omitempty"`
	AppId   int    `json:"app_id,omitempty"`
	OwnerId int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

func New() *CallbackEvent {
	return &CallbackEvent{}
}

func (e *CallbackEvent) SetType(t string) error {
	if _, ok := allowedTypes[t]; !ok {
		return errors.New("incorrect action type")
	}
	e.Type = t
	return nil
}

func (e *CallbackEvent) SetText(text string) {
	e.Text = text
}

func (e *CallbackEvent) SetLink(link string) {
	e.Link = link
}

func (e *CallbackEvent) ToEventData() (string, error) {
	eventJSON, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(eventJSON), nil
}
