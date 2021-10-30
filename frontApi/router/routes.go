package router

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	r.GET("/servers", controllers.GetServers)
}