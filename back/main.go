package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/discord"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/helpers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/api/routes"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Init db connection
	err := database.Init()
	helpers.DieOnError("Database connection failed", err)
	database.Migrate()

	// Init and launch router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			env.GetVariable("SERVER_ADDR_FRONT"),
		},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Authorization"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "User-Agent", "Referrer", "Host"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.Initialize(router)

	infrastructure.InitCommandBus()
	if _, err := discord.InitializeBot(); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := router.Run(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// ----------------- CLOSE APP -----------------
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	consume(quit)
	log.Info("Shutting down server...")
}

func consume(ch <-chan os.Signal) os.Signal {
	return <-ch
}