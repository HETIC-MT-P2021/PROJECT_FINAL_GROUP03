package infrastructure

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/bot/domain"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/cqrs"
	log "github.com/sirupsen/logrus"
)

var CommandBus cqrs.CommandBus

func InitCommandBus() {
	log.Info("Initializing command bus")
	CommandBus = *cqrs.NewCommandBus()

	if err := CommandBus.RegisterHandler(&domain.SendInterfaceLinkCommandHandler{}, &domain.SendInterfaceLinkCommand{}); err != nil {
		log.Error("cannot register handler : ", err)
	}
	if err := CommandBus.RegisterHandler(&domain.ChangeWelcomeMessageCommandHandler{}, &domain.ChangeWelcomeMessageCommand{}); err != nil {
		log.Error("cannot register handler : ", err)
	}
}
