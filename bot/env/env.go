package env

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

// GetVariable loads a variable from the .env file
func GetVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Error("Error loading .env file", err)
	}

	envVariable, variableExists := os.LookupEnv(key)
	if !variableExists {
		log.Error("Couldn't find variable : ", key)
	}

	return envVariable
}
