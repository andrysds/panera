package app

import (
	"fmt"

	"github.com/andrysds/panera/config"
	"github.com/andrysds/panera/handler"
	"gopkg.in/telegram-bot-api.v4"
)

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run() {
	var command string
	for {
		fmt.Print("> ")
		fmt.Scan(&command)
		if command == "exit" {
			return
		}
		c.Handle(command)
	}
}

func (c *Cli) Handle(command string) {
	var message *tgbotapi.MessageConfig
	handler.LogMessage("input", command)
	message = handler.HandleCommand(config.MasterID, command, nil)
	c.SendMessage(message)
}

func (c *Cli) SendMessage(message *tgbotapi.MessageConfig) {
	if message != nil {
		handler.LogMessage("output", message.Text)
	}
}
