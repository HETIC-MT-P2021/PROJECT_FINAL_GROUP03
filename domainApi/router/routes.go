package router

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	r.POST("/commands/change-welcome-message", controllers.ChangeWelcomeMessage)
}