package bot

import (
	"log"

	"github.com/lex3man/playground/cmd/bot/handlers"
	"github.com/lex3man/playground/internal/models"
	statemanagment "github.com/lex3man/playground/internal/utils/stateManagment"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(ch chan models.Task, states *statemanagment.StateRepo, users *map[int]*models.User) {
	bot, err := tgbotapi.NewBotAPI("7648853295:AAHlSwwmpSMPk9f8pltEyo1N9jGt2YgiFuY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.Text[0] == '/' {
				handlers.CommandRouter(bot, &update, states, users)
			} else {
				handlers.MessageRouter(bot, &update, states, users)
			}
		}
	}
}
