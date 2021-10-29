package router

import (
	"github.com/gin-gonic/gin"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/controllers"
)

func Initialize(r *gin.Engine) {
	r.GET("/servers", controllers.GetServers)
}