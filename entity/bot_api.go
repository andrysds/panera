package entity

import (
	"gopkg.in/telegram-bot-api.v4"
)

type BotAPI interface {
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	KickChatMember(tgbotapi.KickChatMemberConfig) (tgbotapi.APIResponse, error)
	GetInviteLink(tgbotapi.ChatConfig) (string, error)
}

type CliBotAPI struct{}

func NewCliBotAPI() *CliBotAPI {
	return &CliBotAPI{}
}

func (c *CliBotAPI) Send(chattable tgbotapi.Chattable) (tgbotapi.Message, error) {
	if message, ok := chattable.(*tgbotapi.MessageConfig); ok {
		LogMessage("output", message.Text)
	}
	return tgbotapi.Message{}, nil
}

func (c *CliBotAPI) KickChatMember(config tgbotapi.KickChatMemberConfig) (tgbotapi.APIResponse, error) {
	return tgbotapi.APIResponse{}, nil
}

func (c *CliBotAPI) GetInviteLink(config tgbotapi.ChatConfig) (string, error) {
	return "https://t.me/joinchat/abcde", nil
}
