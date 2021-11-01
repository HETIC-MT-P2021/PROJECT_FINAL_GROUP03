package controllers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetServers(c *gin.Context) {
	authCode := c.GetHeader("Authorization")
	if "" == authCode {
		c.JSON(http.StatusUnauthorized,"Authorization code needed")
		return
	}

	session, err := services.GetUserSession(authCode)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			"couldn't connect to discord, please check the authorization code or try again later",
		)
		return
	}

	// Get all user guilds
	userGuilds, err := session.UserGuilds(100, "", "")
	if err != nil {
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
				log.Info("ICONE ", userGuild.Icon)

				finalServers = append(finalServers, guild)
			}
		}
	}

	c.JSON(http.StatusOK, finalServers)
}