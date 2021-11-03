package services

import (
	"encoding/json"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/env"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetBotServers() ([]models.Server, error) {
	servers :=  make([]models.Server, 0)
	client := &http.Client{}

	r, err := http.NewRequest("GET", env.GetVariable("DOMAIN_API_URL") + "/servers", strings.NewReader(""))
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

	log.Info(string(body))
	json.Unmarshal(body, &servers)

	return servers, err
}