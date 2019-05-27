package connection

import (
	"log"
	"os"
	"strconv"

	"github.com/andrysds/clarity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// Telegram represents Telegram connection
	Telegram *tgbotapi.BotAPI

	// MasterTelegramID stores master ID from env
	MasterTelegramID int64
	// SquadTelegramID stores squad ID from env
	SquadTelegramID int64
	// BirthdayTelegramID stores birthday ID from env
	BirthdayTelegramID int64
)

// InitTelegram initializes Telegram connection
func InitTelegram() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		return
	}
	MasterTelegramID, _ = strconv.ParseInt(os.Getenv("MASTER_ID"), 10, 64)
	SquadTelegramID, _ = strconv.ParseInt(os.Getenv("SQUAD_ID"), 10, 64)
	BirthdayTelegramID, _ = strconv.ParseInt(os.Getenv("BIRTHDAY_ID"), 10, 64)

	var err error
	Telegram, err = tgbotapi.NewBotAPI(token)
	clarity.PanicIfError(err, "error on creating telegram bot api connection")
	log.Printf("* Authorized on account %s\n", Telegram.Self.UserName)

	webhook := tgbotapi.NewWebhook(os.Getenv("WEBHOOK_URL") + Telegram.Token)
	_, err = Telegram.SetWebhook(webhook)
	clarity.PanicIfError(err, "error on setting bot webhook")
}
