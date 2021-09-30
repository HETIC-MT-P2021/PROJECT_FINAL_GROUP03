package services

import (
	"errors"
	"net/http"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/domain"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	"github.com/gin-gonic/gin"
)

func GetAccountFromContext(c *gin.Context) (models.Account, error) {
	payload, ok := c.Get("account")
	if !ok {
		return models.Account{}, errors.New("no account found")
	}

	return payload.(models.Account), nil
}

// HasRightOnServerFromContext returns true if account found in context has right on server found in param
func HasRightOnServerFromContext(c *gin.Context) bool {
	account, err := GetAccountFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "nope")
		return false
	}

	return domain.Account.IsMember(account.DiscordID, c.Param("id"))
}
