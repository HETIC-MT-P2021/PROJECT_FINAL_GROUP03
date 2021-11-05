package controllers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/services"
	"github.com/bwmarrin/discordgo"
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
				isNotAdmin, err := services.MemberHasPermission(session, guild.DiscordID, discordgo.PermissionAdministrator)
				if err != nil {
					log.Error(err)
				}

				if !userGuild.Owner && isNotAdmin {
					continue
				}

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
		DiscordID:       serverID,
		Name:            userGuild.Name,
		WelcomeMessage:  botGuild.WelcomeMessage,
		BirthdayMessage: botGuild.BirthdayMessage,
		ForbiddenWords:  botGuild.ForbiddenWords,
	})
}

func PatchServer(c *gin.Context) {
	serverID := c.Param("id")
	var payload models.Server
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, "payload should be compatible with models.Server struct")
		return
	}

	if payload.WelcomeMessage != "" {
		if err := services.ChangeWelcomeMessage(serverID, payload.WelcomeMessage); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	if payload.ForbiddenWords != "" {
		if err := services.ChangeForbiddenWords(serverID, payload.ForbiddenWords); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	if payload.BirthdayMessage != "" {
		if err := services.ChangeBirthdayMessage(serverID, payload.BirthdayMessage); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, "")
}
