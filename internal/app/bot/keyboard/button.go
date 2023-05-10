package keyboard

import (
	"errors"
)

var (
	allowedTypes = map[string]struct{}{
		"text":      {},
		"open_link": {},
		"location":  {},
		"vkpay":     {},
		"open_app":  {},
		"callback":  {},
	}
	allowedColors = map[string]struct{}{
		"primary":   {},
		"secondary": {},
		"negative":  {},
		"positive":  {},
	}
)

type Button struct {
	Action struct {
		Type    string `json:"type"`
		Label   string `json:"label,omitempty"`
		Payload string `json:"payload,omitempty"`
		Link    string `json:"link,omitempty"`
		Hash    string `json:"hash,omitempty"`
		AppId   int    `json:"app_id,omitempty"`
		OwnerId int    `json:"owner_id,omitempty"`
	} `json:"action"`
	Color string `json:"color,omitempty"`
}

func NewButton() *Button {
	return &Button{}
}

func (b *Button) SetColor(color string) error {
	if _, ok := allowedColors[color]; !ok {
		return errors.New("incorrect button color")
	}
	b.Color = color
	return nil
}

func (b *Button) SetType(t string) error {
	if _, ok := allowedTypes[t]; !ok {
		return errors.New("incorrect action type")
	}
	b.Action.Type = t
	return nil
}

func (b *Button) SetLabel(label string) {
	b.Action.Label = label
}

func (b *Button) SetPayload(pl string) {
	b.Action.Payload = pl
}

func (b *Button) SetLink(link string) {
	b.Action.Link = link
}
