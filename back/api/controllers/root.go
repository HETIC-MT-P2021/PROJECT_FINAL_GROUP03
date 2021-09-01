package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}