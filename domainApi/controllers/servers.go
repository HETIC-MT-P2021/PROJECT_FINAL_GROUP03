package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeWelcomeMessage(c *gin.Context) {
	var server models.Server
	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain a DiscordId and a WelcomeMessage")
		return
	}

	foundServer := models.Server{
		DiscordID: server.DiscordID,
	}
	if err := repositories.FindServerByDiscordID(&foundServer); err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Server with id %s not found", server.DiscordID))
		return
	}

	if err := repositories.UpdateServerMessage(&server); err != nil {
		c.JSON(http.StatusInternalServerError, "couldnt update server message, please try again or contact support")
		return
	}

	c.JSON(http.StatusOK, "server message updated successfully")
}