package controllers

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if "" == accessToken {
		c.JSON(http.StatusUnauthorized,"Authorization code needed")
		return
	}

	session, err := services.GetUserSession(accessToken)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			"couldn't connect to discord, please check the authorization code or try again later",
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
