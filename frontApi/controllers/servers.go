package controllers

import "github.com/gin-gonic/gin"

func GetServers(c *gin.Context) {
	c.JSON(200, "hey servers")
}