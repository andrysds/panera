package config

import (
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	BotToken string

	Port       string
	WebhookUrl string

	RedisUrl string

	MasterID int64
	SquadID  int64

	Username string
	Password string
)

func Init() {
	gotenv.Load()

	BotToken = os.Getenv("BOT_TOKEN")

	Port = os.Getenv("PORT")
	WebhookUrl = os.Getenv("WEBHOOK_URL")

	RedisUrl = os.Getenv("REDIS_URL")

	MasterID, _ = strconv.ParseInt(os.Getenv("MASTER_ID"), 10, 64)
	SquadID, _ = strconv.ParseInt(os.Getenv("SQUAD_ID"), 10, 64)

	Username = os.Getenv("USERNAME")
	Password = os.Getenv("PASSWORD")

	log.Println("* Config initialized")
}
