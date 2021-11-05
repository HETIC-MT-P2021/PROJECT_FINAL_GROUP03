package controllers

import (
	"fmt"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ChangeWelcomeMessage(c *gin.Context) {
	var server models.Server
	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain a discord_id and a welcome_message")
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

func ChangeBirthdayMessage(c *gin.Context) {
	var server models.Server
	if err := c.ShouldBindJSON(&server); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, "request should contain a DiscordId and a Birthday message")
		return
	}
	log.Info(server)

	foundServer := models.Server{
		DiscordID: server.DiscordID,
	}
	if err := repositories.FindServerByDiscordID(&foundServer); err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Server with id %s not found", server.DiscordID))
		return
	}

	if err := repositories.UpdateBirthdayMessage(&server); err != nil {
		c.JSON(http.StatusInternalServerError, "couldnt update server message, please try again or contact support")
		return
	}

	c.JSON(http.StatusOK, "Birthday message updated successfully")
}

func ChangeForbiddenWordsList(c *gin.Context) {
	var server models.Server

	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, "request should contain a discord_id and a forbidden_words")
		return
	}

	foundServer := models.Server{
		DiscordID: server.DiscordID,
	}

	if err := repositories.FindServerByDiscordID(&foundServer); err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Server with id %s not found", server.DiscordID))
		return
	}

	if err := repositories.UpdateServerForbiddenWords(&server); err != nil {
		c.JSON(http.StatusInternalServerError, "couldnt update server words, please try again or contact support")
		return
	}

	c.JSON(http.StatusOK, "server forbidden words updated successfully")
}

func GetAll(c *gin.Context) {
	var servers []models.Server

	if err := repositories.FindAllServers(&servers); err != nil {
		c.JSON(http.StatusInternalServerError, "error while fetching servers")
		return
	}

	c.JSON(http.StatusOK, servers)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	server := models.Server{
		DiscordID: id,
	}

	if err := repositories.FindServerByDiscordID(&server); err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, server)
}
