package main

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/handlers"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/JackMaarek/go-bot-utils/connectors"
)

func main() {
	// init bot connection
	dg, err := connectors.InitializeBot()
	dg.AddHandler(handlers.MessageCreate)

	if err != nil {
		log.Fatal(err)
	}
	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
