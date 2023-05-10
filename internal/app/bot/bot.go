package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"vk_go_develop_test_task/internal/app/bot/callbackEvent"
	"vk_go_develop_test_task/internal/app/bot/keyboard"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"

	"vk_go_develop_test_task/internal/app/config"
)

type Bot struct {
	vk    *api.VK
	group api.GroupsGetByIDResponse
	lp    *longpoll.LongPoll
}

func New(cfg *config.Config) (*Bot, error) {
	b := &Bot{}
	var err error

	b.vk = api.NewVK(cfg.BotToken)
	// Получаем информацию о группе
	b.group, err = b.vk.GroupsGetByID(api.Params{})
	if err != nil {
		return nil, err
	}

	// Инициализируем longpoll
	b.lp, err = longpoll.NewLongPoll(b.vk, b.group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	b.lp.MessageNew(b.NewMessageHandler)
	b.lp.MessageEvent(b.MessageEventHandler)

	return b, nil
}

func (b *Bot) Run() error {
	err := b.lp.Run()
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) NewMessageHandler(ctx context.Context, message events.MessageNewObject) {
	p := params.NewMessagesSendBuilder()
	p.PeerID(message.Message.FromID)
	p.RandomID(0)
	p.Message(message.Message.Text)

	k := keyboard.New()

	k.SetInline(false)
	k.SetOneTime(false)

	b1, _ := createTextButton("red", "negative")
	b2, _ := createCallbackButton("first callback", "{\"button\": \"1\"}")
	b3, _ := createCallbackButton("second callback", "{\"button\": \"2\"}")
	b4, _ := createCallbackButton("third callback", "{\"button\": \"3\"}")

	row1 := []keyboard.Button{*b1, *b2, *b3, *b4}
	k.AddButtonRow(row1)

	b5, _ := createOpenLinkButton("Link", "https://youtu.be/dQw4w9WgXcQ")
	b6, _ := createLocationButton()

	row2 := []keyboard.Button{*b5, *b6}
	k.AddButtonRow(row2)

	kb, err := k.ToParams()
	if err != nil {
		log.Println(err)
	}

	p.Keyboard(kb)

	resp, err := b.vk.MessagesSend(p.Params)
	if err != nil {
		log.Println(fmt.Errorf("send message error, response: %d error: %e", resp, err))
	}
}

func (b *Bot) MessageEventHandler(ctx context.Context, message events.MessageEventObject) {

	p := params.NewMessagesSendMessageEventAnswerBuilder()
	p.PeerID(message.PeerID)
	p.UserID(message.UserID)
	p.EventID(message.EventID)

	type buttonNum struct {
		Button string `json:"button"`
	}

	var but buttonNum
	err := json.Unmarshal(message.Payload, &but)
	if err != nil {
		log.Println(fmt.Errorf("unmarshal payload error: %e", err))
		return
	}

	var e *callbackEvent.CallbackEvent
	switch but.Button {
	case "1":
		e, err = firstCallbackButtonEvent()
		if err != nil {
			log.Println(fmt.Errorf("first button callback error: %e", err))
			return
		}
	case "2":
		e, err = secondCallbackButtonEvent()
		if err != nil {
			log.Println(fmt.Errorf("second button callback error: %e", err))
			return
		}
	case "3":
		e, err = thirdCallbackButtonEvent()
		if err != nil {
			log.Println(fmt.Errorf("thrid button callback error: %e", err))
			return
		}
	}

	event, err := e.ToEventData()
	if err != nil {
		log.Println(fmt.Errorf("marshal to JSON error: %e", err))
		return
	}
	p.EventData(event)

	resp, err := b.vk.MessagesSendMessageEventAnswer(p.Params)
	if err != nil {
		log.Println(fmt.Errorf("send message error, response: %d error: %e", resp, err))
	}
}

func firstCallbackButtonEvent() (*callbackEvent.CallbackEvent, error) {
	e := callbackEvent.New()
	err := e.SetType("show_snackbar")
	if err != nil {
		return nil, err
	}

	e.SetText("Исчезающее сообщение")

	return e, nil
}

func secondCallbackButtonEvent() (*callbackEvent.CallbackEvent, error) {
	e := callbackEvent.New()
	err := e.SetType("open_link")
	if err != nil {
		return nil, err
	}

	e.SetLink("https://vk.com/id0")

	return e, nil
}

func thirdCallbackButtonEvent() (*callbackEvent.CallbackEvent, error) {
	e := callbackEvent.New()
	err := e.SetType("show_snackbar")
	if err != nil {
		return nil, err
	}

	e.SetText("Другое исчезающее сообщение")

	return e, nil
}

func createTextButton(text string, color string) (*keyboard.Button, error) {
	b := keyboard.NewButton()
	err := b.SetType("text")
	if err != nil {
		return nil, err
	}

	b.SetLabel(text)
	err = b.SetColor(color)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func createOpenLinkButton(text, link string) (*keyboard.Button, error) {
	b := keyboard.NewButton()
	err := b.SetType("open_link")
	if err != nil {
		return nil, err
	}

	b.SetLabel(text)
	b.SetLink(link)

	return b, nil
}

func createLocationButton() (*keyboard.Button, error) {
	b := keyboard.NewButton()
	err := b.SetType("location")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func createCallbackButton(label, payload string) (*keyboard.Button, error) {
	b := keyboard.NewButton()
	err := b.SetType("callback")
	if err != nil {
		return nil, err
	}

	b.SetLabel(label)
	b.SetPayload(payload)

	return b, nil
}
