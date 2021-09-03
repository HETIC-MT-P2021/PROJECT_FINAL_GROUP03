package database

import (
	"fmt"
	"time"

	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/back/shared/helpers"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"errors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db is the database object
var Db *gorm.DB

// Config is the structure used to load db credentials from the environment.
type Config struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     uint64 `env:"DB_PORT" envDefault:"5432"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
}

// Init Initializes a db connection
func Init() error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	dbURL := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	var tmpDb *gorm.DB

	// Try connecting database 5 times
	for test := 1; test <= 5; test++ {
		tmpDb, err = gorm.Open("postgres", dbURL)

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
		User:     env.GetVariable("DB_USER"),
		Password: env.GetVariable("DB_PASSWORD"),
		Port:     helpers.ParseStringToUint64(env.GetVariable("DB_PORT")),
		Name:     env.GetVariable("DB_NAME"),
		Host:     env.GetVariable("DB_HOST"),
	}
	if cfg.User == "" ||
		cfg.Password == "" ||
		cfg.Host == "" ||
		cfg.Name == "" {
		err = errors.New("Missing env variables")
	}

	return cfg, err
}
