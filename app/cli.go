package app

import (
	"fmt"

	"github.com/andrysds/panera/entity"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

type Cli struct {
	BotAPI entity.BotAPI
}

func NewCli() *Cli {
	return &Cli{
		BotAPI: entity.NewCliBotAPI(),
	}
}

func (c *Cli) GetBotAPI() entity.BotAPI {
	return c.BotAPI
}

func (c *Cli) Run() {
	var command string
	update := entity.BlankUpdate
	for {
		fmt.Print("> ")
		fmt.Scan(&command)
		if command == "exit" {
			return
		}
		update.Message.Text = "/" + command
		update.Message.From.UserName = "master"
		update.Message.Entities = &[]tgbotapi.MessageEntity{
			tgbotapi.MessageEntity{
				Type:   "bot_command",
				Offset: 0,
				Length: len(update.Message.Text),
			},
		}
		handler.HandleUpdate(update, c.BotAPI)
	}
}
