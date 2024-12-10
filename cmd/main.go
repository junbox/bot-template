package main

import (
	"github.com/lex3man/playground/cmd/bot"
	"github.com/lex3man/playground/cmd/webhook"
	"github.com/lex3man/playground/internal/models"
	sm "github.com/lex3man/playground/internal/utils/stateManagment"
)

func main() {
	TOKEN = "???"
	users := make(map[int]*models.User)
	states := sm.StateRepo{
		States: make(map[int]*sm.UserState),
		Vars:   make(map[string]map[int]string),
	}
	states.States[0] = &sm.UserState{}
	ch := make(chan models.Task)
	go webhook.StartHook(ch, &states, &users)
	bot.StartBot(ch, &states, &users, TOKEN)
}
