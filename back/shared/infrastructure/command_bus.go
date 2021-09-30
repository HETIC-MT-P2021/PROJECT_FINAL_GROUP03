package infrastructure

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/commands"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	log "github.com/sirupsen/logrus"
)

var CommandBus cqrs.CommandBus

func InitCommandBus() {
	log.Info("Initializing command bus")
	CommandBus = *cqrs.NewCommandBus()

	if err := CommandBus.RegisterHandler(&commands.SendInterfaceLinkCommandHandler{}, &commands.SendInterfaceLinkCommand{}); err != nil {
		log.Error("cannot register handler : ", err)
	}
	if err := CommandBus.RegisterHandler(&commands.ChangeWelcomeMessageCommandHandler{}, &commands.ChangeWelcomeMessageCommand{}); err != nil {
		log.Error("cannot register handler : ", err)
	}
}
