package webhook

import (
	"log"
	"net/http"

	"github.com/lex3man/playground/internal/models"
	statemanagment "github.com/lex3man/playground/internal/utils/stateManagment"

	"github.com/gin-gonic/gin"
)

func StartHook(ch chan models.Task, states *statemanagment.StateRepo, users *map[int]*models.User) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		log.Panic("Bot not started!!!")
	}
}
