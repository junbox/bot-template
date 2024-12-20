package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lex3man/playground/internal/models"
	statemanagment "github.com/lex3man/playground/internal/utils/stateManagment"
)

func CommandRouter(bot *tgbotapi.BotAPI, update *tgbotapi.Update, states *statemanagment.StateRepo, users *map[int]*models.User) {
	var kb tgbotapi.ReplyKeyboardMarkup
	var msgText string
	isRemoveKeyboard := tgbotapi.NewRemoveKeyboard(true)

	switch update.Message.Text {
	case "/start":
		isRemoveKeyboard.Selective = false
		kb = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Привет"),
				tgbotapi.NewKeyboardButton("Пока!"),
			),
		)
		states.SetDefault(int(update.Message.From.ID))
		msgText = "Чего тебе?"
	default:
		isRemoveKeyboard.Selective = true
		msgText = "Не знаю такую команду"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
	// msg.ReplyToMessageID = update.Message.MessageID
	if !isRemoveKeyboard.Selective {
		msg.ReplyMarkup = kb
	} else {
		msg.ReplyMarkup = isRemoveKeyboard
	}

	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Message not sent")
	}
}
