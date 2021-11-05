package services

import (
	"encoding/json"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func GetBotServerById(id string) (models.Server, error) {
	var server models.Server
	client := &http.Client{}

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

func ChangeWelcomeMessage(serverID, welcomeMessage string) error {
	payload, err := json.Marshal(models.ChangeWelcomeMessageForm{WelcomeMessage: welcomeMessage, DiscordID: serverID})
	if err != nil {
		return err
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", env.GetVariable("DOMAIN_API_URL")+"/commands/change-welcome-message", strings.NewReader(string(payload)))
	if err != nil {
		return err
	}

	_, err = client.Do(r)

	return err
}
