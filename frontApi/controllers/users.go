package controllers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	session, err := services.GetUserSession(c)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			err.Error(),
		)
		return
	}

	user, err := session.User("@me")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
