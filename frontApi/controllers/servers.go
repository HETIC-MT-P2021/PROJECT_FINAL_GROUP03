package controllers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetServers(c *gin.Context) {
	session, err := services.GetUserSession(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			err.Error(),
		)
		return
	}

	// Get all user guilds
	userGuilds, err := session.UserGuilds(100, "", "")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, "couldn't fetch guilds")
		return
	}

	// Get bot guilds
	botGuilds, err := services.GetBotServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't fetch bot guilds")
		log.Error(err)
		return
	}

	finalServers := make([]models.Server, 0)
	for _, guild := range botGuilds {
		for _, userGuild := range userGuilds {
			if userGuild.ID == guild.DiscordID {
				guild.Name = userGuild.Name

				finalServers = append(finalServers, guild)
			}
		}
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, finalServers)
}

func GetServerByID(c *gin.Context) {

	serverID := c.Param("id")

	session, err := services.GetUserSession(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			err.Error(),
		)
		return
	}

	userGuild, err := services.GetUserGuildByID(session, serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	botGuild, err := services.GetBotServerById(serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, models.Server{
		DiscordID: serverID,
		Name: userGuild.Name,
		WelcomeMessage: botGuild.WelcomeMessage,
	})
}