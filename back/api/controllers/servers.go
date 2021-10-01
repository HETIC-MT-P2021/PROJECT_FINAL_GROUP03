package controllers

import (
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/services"
	"github.com/gin-gonic/gin"
)

func GetServers(c *gin.Context) {
	account, err := services.GetAccountFromContext(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Couldn't find user from provided informations")
		return
	}
	servers, err := repositories.FindAccountServers(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldnt fetch servers")
		return
	}

	c.JSON(http.StatusOK, servers)
}

func GetServerByID(c *gin.Context) {
	if !services.HasRightOnServerFromContext(c) {
		c.JSON(http.StatusUnauthorized, "you are not allowed to do this")
		return
	}
	server := models.Server {
		DiscordID: c.Param("id"),
	}
	if err := repositories.FindServerByDiscordID(&server); err != nil {
		c.JSON(http.StatusInternalServerError, "couldnt fetch server")
		return
	}

	c.JSON(http.StatusOK, server)
}