package statemanagment

import (
	"fmt"
	"strings"

	"github.com/lex3man/playground/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetReply(bot *tgbotapi.BotAPI, update *tgbotapi.Update, states *StateRepo, users *map[int]*models.User) tgbotapi.MessageConfig {
	userID := int(update.Message.From.ID)
	userReply := update.Message.Text
	var kb tgbotapi.ReplyKeyboardMarkup
	isRemoveKeyboard := tgbotapi.NewRemoveKeyboard(true)

	var msgReply string

	switch states.States[userID].State {
	case "registration":
		switch states.States[userID].Step {
		case "name":
			states.SetVar(userID, "name", userReply)
			states.SetState(userID, "registration", "age")
			msgReply = fmt.Sprintf("Хорошо, %s, а сколько тебе лет?", userReply)
		case "age":
			states.SetVar(userID, "age", userReply)
			states.SetState(userID, "registration", "gender")
			msgReply = "Выбери свой пол"
			isRemoveKeyboard.Selective = false
			kb = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Мужской"),
					tgbotapi.NewKeyboardButton("Женский"),
				),
			)
		case "gender":
			states.SetVar(userID, "gender", userReply)
			states.SetState(userID, "registration", "city")
			msgReply = "Из какого ты города?"
		case "city":
			states.SetVar(userID, "city", userReply)
			states.SetState(userID, "registration", "finish")
			msgReply = "Пожалуй этого пока хватит для регистрации"
			isRemoveKeyboard.Selective = false
			kb = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Завершить"),
				),
			)
		case "finish":
			states.SetDefault(userID)
			msgReply = "Я тебя запомнил!"
			user := *(*users)[userID]
			user.Name = states.GetVar(userID, "name")
			user.City = states.GetVar(userID, "city")
			fmt.Println(user)
			user.AddAchivment("HAPPY REGISTRATION")
		default:
			msgReply = "???"
		}
	case "default":
		switch strings.ToLower(userReply) {
		case "зарегистрироваться":
			msgReply = "Давай начнём с того, как ты бы хотел, чтобы я к тебе обращался?"
			states.States[userID].State = "registration"
			states.States[userID].Step = "name"
		case "привет":
			msgReply = "Ну привет! Давай я расскажу тебе..."
			isRemoveKeyboard.Selective = false
			kb = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Зарегистрироваться"),
				),
			)
		case "пока!":
			msgReply = "Может я всё таки могу тебе как-то помочь?"
			isRemoveKeyboard.Selective = false
			kb = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Buy"),
					tgbotapi.NewKeyboardButton("Sell"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Not now"),
				),
			)
		default:
			msgReply = "Я пока не всё умею"
		}
	default:
		msgReply = "???"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgReply)
	if isRemoveKeyboard.Selective {
		msg.ReplyMarkup = isRemoveKeyboard
	} else {
		msg.ReplyMarkup = kb
	}

	return msg
}
