package cron

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/andrysds/panera/connection"
	"github.com/andrysds/panera/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
)

// Cron represents application cron jobs
var Cron *cron.Cron

// Init inits cron jobs
func Init() {
	Cron = cron.New()

	Cron.AddFunc("0 30 13 * * 1-5", standupJob)
	Cron.AddFunc("0 0 5 * * 1-5", newDayStandupJob)

	go Cron.Start()
	log.Println("* Cron initialized")
}

func standupJob() {
	if !isHoliday() {
		msg := tgbotapi.NewMessage(connection.SquadTelegramID, handler.Command("standup"))
		connection.Telegram.Send(msg)
	}
}

func newDayStandupJob() {
	if !isHoliday() {
		msg := tgbotapi.NewMessage(connection.MasterTelegramID, handler.Command("standup_new_day"))
		connection.Telegram.Send(msg)
	}
}

func isHoliday() bool {
	holidays := os.Getenv("HOLIDAYS")
	day := time.Now().Format("02")
	return strings.Contains(holidays, day)
}
