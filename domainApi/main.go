package main

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/cron"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/handlers"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/connectors"
	"github.com/JackMaarek/go-bot-utils/database"
)

func main() {
	// init db connection
	if err := database.Init(); err != nil {
		log.Fatal("Could not connect to db : ", err)
	}
	database.Migrate()

	// connect to discord
	session, err := connectors.InitializeBot()
	if err != nil {
		log.Fatal("Could not connect to discord : ", err)
	}
	session.AddHandler(handlers.ForbiddenMessageHandler)
	session.AddHandler(handlers.GuildMemberAdd)

	// Cron jobs
	cron.InitBirthdayReminderJob(session)

	// init router
	r := gin.Default()
	router.Initialize(r)
	router.Run(r)

	log.Info("Hey c'est l'API domain")
}
