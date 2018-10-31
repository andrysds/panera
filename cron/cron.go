package cron

import (
	"log"

	"github.com/andrysds/panera/entity"
	"github.com/robfig/cron"
	"gopkg.in/telegram-bot-api.v4"
)

var Cron *cron.Cron

func Init() {
	Cron = cron.New()

	AddStandupJobs()
	AddBirthdayJobs()

	go Cron.Start()
	log.Println("* Cron initialized")
}

func UpdateFromCron(chatId int64) *tgbotapi.Update {
	update := entity.BlankUpdate
	update.Message.Chat.ID = chatId
	return update
}
