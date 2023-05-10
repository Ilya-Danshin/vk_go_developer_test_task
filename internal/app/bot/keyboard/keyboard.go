package keyboard

import "encoding/json"

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Buttons [][]Button `json:"buttons"`
	Inline  bool       `json:"inline,omitempty"`
}

func New() *Keyboard {
	return &Keyboard{}
}

func (k *Keyboard) AddButtonRow(buttons []Button) {
	k.Buttons = append(k.Buttons, buttons)
}

func (k *Keyboard) SetOneTime(oneTime bool) {
	k.OneTime = oneTime
}

func (k *Keyboard) SetInline(inline bool) {
	k.Inline = inline
}

func (k *Keyboard) ToParams() (string, error) {
	keyboardJSON, err := json.Marshal(k)
	if err != nil {
		return "", err
	}

	return string(keyboardJSON), nil
}
