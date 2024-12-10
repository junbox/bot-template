package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lex3man/playground/internal/models"
	statemanagment "github.com/lex3man/playground/internal/utils/stateManagment"
)

func MessageRouter(bot *tgbotapi.BotAPI, update *tgbotapi.Update, states *statemanagment.StateRepo, users *map[int]*models.User) {
	msg := statemanagment.GetReply(bot, update, states)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Message not sent")
	}
}
