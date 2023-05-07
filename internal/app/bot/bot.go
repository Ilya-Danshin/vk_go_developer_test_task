package bot

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
	"vk_go_develop_test_task/internal/app/config"
)

type Bot struct {
	vk *api.VK
	lp *longpoll.LongPoll
}

func New(cfg *config.Config) (*Bot, error) {
	b := &Bot{}

	b.vk = api.NewVK(cfg.BotToken)
	// Получаем информацию о группе
	group, err := b.vk.GroupsGetByID(api.Params{})
	if err != nil {
		return nil, err
	}

	// Инициализируем longpoll
	b.lp, err = longpoll.NewLongPoll(b.vk, group[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	return b, nil
}

func (b *Bot) Run() error {
	return nil
}
