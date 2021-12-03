package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/JackMaarek/go-bot-utils/enum"
	"github.com/JackMaarek/go-bot-utils/env"
	"github.com/JackMaarek/go-bot-utils/helpers"
	"github.com/JackMaarek/go-bot-utils/models"
)

// GetBotServers makes an http call to the domain api to get bot available servers
func GetBotServers() ([]models.Server, error) {
	servers := make([]models.Server, 0)
	client := &http.Client{}

	r, err := http.NewRequest("GET", env.GetVariable("DOMAIN_API_URL")+"/servers", strings.NewReader(""))
	if err != nil {
		return servers, err
	}

	res, err := client.Do(r)
	if err != nil {
		return servers, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return servers, err
	}

	err = json.Unmarshal(body, &servers)

	return servers, err
}

// GetBotServerById retrieves a bot server by its id
func GetBotServerById(id string) (models.Server, error) {
	var server models.Server
	client := &http.Client{}
	log.Info(env.GetVariable("DOMAIN_API_URL"))

	r, err := http.NewRequest("GET", env.GetVariable("DOMAIN_API_URL")+"/servers/"+id, strings.NewReader(""))
	if err != nil {
		return server, err
	}

	res, err := client.Do(r)
	if err != nil {
		return server, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return server, err
	}

	err = json.Unmarshal(body, &server)

	return server, err
}

// ChangeWelcomeMessage makes an http call to the domain api to change the welcome message for a given server
func ChangeWelcomeMessage(serverID, welcomeMessage string) error {
	payload, err := json.Marshal(models.ChangeWelcomeMessageForm{WelcomeMessage: welcomeMessage, DiscordID: serverID})
	if err != nil {
		return err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", env.GetVariable("DOMAIN_API_URL")+"/cmd/change-welcome-message", strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	_, err = client.Do(r)

	return err
}

// ChangeBirthdayMessage makes an http call to the domain api to change the borthday message of a server
func ChangeBirthdayMessage(serverID string, birthdayMessage string) error {
	response, err := helpers.PerformRequest(enum.ChangeBirthdayMessageRoute, enum.Post, models.ChangeBirthdayMessage{
		BirthdayMessage: birthdayMessage,
		DiscordID:       serverID,
	})
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		log.Error("Error occurred while updating birthday message")
		return err
	}

	return nil
}

// ChangeForbiddenWords makes an http call to the domain api to change the forbidden words list for a server
func ChangeForbiddenWords(serverID, wordsList string) error {
	payload, err := json.Marshal(models.ForbiddenWordsListForm{ForbiddenWords: wordsList, DiscordID: serverID})
	if err != nil {
		return err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", env.GetVariable("DOMAIN_API_URL")+"/cmd/change-forbidden-words", strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	_, err = client.Do(r)

	return err
}
