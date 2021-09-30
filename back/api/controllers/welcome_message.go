package controllers

import (
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UpdateWelcomeMessage(c *gin.Context) {
	if !services.HasRightOnServerFromContext(c) {
		c.JSON(http.StatusUnauthorized, "you are not allowed to do this")
		return
	}
	var form WelcomeMessageForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, "new welcome message needed")
		return
	}
	log.Info("new message : ", form.WelcomeMessage)
	cmd := cqrs.NewCommandMessage(&commands.ChangeWelcomeMessageCommand{
		Session:         nil,
		ServerDiscordID: c.Param("id"),
		WelcomeMessage:  form.WelcomeMessage,
	})

	_, err := infrastructure.CommandBus.Dispatch(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't execute your demand")
		return
	}

	c.JSON(http.StatusOK, "message updated with success")
}

type WelcomeMessageForm struct {
	WelcomeMessage string `json:"welcome-message"`
}
