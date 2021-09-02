package routes

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/api/controllers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Initialize(r *gin.Engine) {
	configureCORS(r)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/", controllers.SayHello)
}

func configureCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			env.GetVariable("SERVER_ADDR_FRONT"),
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Authorization"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "User-Agent", "Referrer", "Host"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
