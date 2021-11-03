package database

import (
	"time"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db is the database object
var Db *gorm.DB

// Config is the structure used to load db credentials from the environment.
type Config struct {
	DbURL    string `env:"DATABASE_URL"`
}

// Init Initializes a db connection
func Init() error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	var tmpDb *gorm.DB
	// Try connecting database 5 times
	for test := 1; test <= 5; test++ {
		tmpDb, err = gorm.Open(postgres.Open(cfg.DbURL), &gorm.Config{})

		if err != nil {
			log.Warnf("db connection failed. (%d/5)", test)
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}
	if err != nil {
		return err
	}

	Db = tmpDb
	log.Info("Connected to database!")

	return nil
}

func getConfig() (Config, error) {
	var err error
	cfg := Config{
		DbURL:    env.GetVariable("DATABASE_URL"),
	}

	return cfg, err
}
