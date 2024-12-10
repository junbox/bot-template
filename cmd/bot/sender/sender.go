package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lex3man/playground/internal/models"
)

func Send(bot *tgbotapi.BotAPI, ch chan models.Task) {
	for {
		if ch != nil {
			log.Println(ch)
		}
	}
}
