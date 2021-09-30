package database

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	Db.AutoMigrate(&models.Account{}, &models.Server{})
}
