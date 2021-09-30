package controllers

import (
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/api/domain"
	botDomain "github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/domain"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UpdateWelcomeMessage(c *gin.Context) {
	server := models.Server{
		DiscordID: c.Param("id"),
	}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		c.JSON(http.StatusNotFound, "server does not exist")
		return
	}
	payload, ok := c.Get("account")
	if !ok {
		c.JSON(http.StatusUnauthorized, "user not found")
		return
	}
	account := payload.(models.Account)
	if !domain.HasRight(&account, &server) {
		c.JSON(http.StatusUnauthorized, "you cannot edit this server welcome message")
		return
	}
	var form WelcomeMessageForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, "new welcome message needed")
		return
	}
	log.Info("new message : ", form.WelcomeMessage)
	cmd := cqrs.NewCommandMessage(&botDomain.ChangeWelcomeMessageCommand{
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
