package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/discord"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/helpers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/infrastructure"
	"github.com/gin-gonic/gin"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/api/routes"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Init db connection
	err := database.Init()
	helpers.DieOnError("Database connection failed", err)
	database.Migrate()

	// Init and launch router
	router := gin.Default()
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
