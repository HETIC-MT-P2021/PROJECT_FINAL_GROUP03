package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetServers(c *gin.Context) {
	authCode := c.GetHeader("Authorization")
	if "" == authCode {
		c.JSON(http.StatusUnauthorized,"Authorization code needed")
		return
	}

	c.JSON(http.StatusOK, "here are your servers")
}