package database

import (
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/domainApi/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	err := Db.AutoMigrate(&models.Server{})
	if err != nil {
		log.Warnf("Could not execute migrations")
	}
}