package middlewares

import (
	"strings"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CheckAccount(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) < 2 {
		log.Error("invalid hash : ", parts)
		return
	}

	account := models.Account{
		DiscordID: parts[1],
	}
	if err := repositories.FindAccountByDiscordID(&account); err != nil {
		return
	}
	c.Set("account", account)
	c.Next()
}
