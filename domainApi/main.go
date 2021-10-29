package main

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/database"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/discordApi"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/router"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var DiscordSession *discordgo.Session

func main() {
	// init db connection
	if err := database.Init(); err != nil {
		log.Fatal("Could not connect to db : ", err)
	}
	database.Migrate()

	// connect to discord
	var err error
	DiscordSession, err = discordApi.InitializeBot()
	if err != nil {
		log.Fatal("Could not connect to discord : ", err)
	}

	// init router
	r := gin.Default()
	router.Initialize(r)
	router.Run(r)


	log.Info("Hey c'est l'API domain")	
}