package router

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/commands/change-welcome-message", controllers.ChangeWelcomeMessage)
		v1.POST("/commands/change-birthday-message", controllers.ChangeBirthdayMessage)

		v1.POST("/birthdays", controllers.CreateUserBirthday)

		v1.GET("/servers", controllers.GetAll)
		v1.GET("/servers/:id", controllers.GetByID)
	}

}
