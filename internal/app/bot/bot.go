package bot

import (
	"context"
	"fmt"
	"log"

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

	resp, err := b.vk.MessagesSend(p.Params)
	if err != nil {
		log.Println(fmt.Errorf("send message error, response: %d error: %e", resp, err))
	}
}
