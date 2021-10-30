package services

import (
	"encoding/json"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/env"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetUserSession(code string) (*discordgo.Session, error) {
	body, err := getUserToken(code)
	if err != nil {
		log.Fatal(err)
	}
	var decodedBody map[string]interface{}

	if err := json.Unmarshal([]byte(body), &decodedBody); err != nil {
		log.Fatal("could not decode body ", err)
	}
	log.Info(decodedBody["access_token"].(string))

	return discordgo.New("Bearer " + decodedBody["access_token"].(string))
}

func getUserToken(code string) (string, error) {
	clientId := env.GetVariable("CLIENT_ID")
	clientSecret := env.GetVariable("CLIENT_SECRET")

	endpoint := "https://discord.com/api/oauth2/token"
	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("scope", "guilds%20identify")
	data.Set("redirect_uri", "http://localhost:8080")
	data.Set("grant_type", "authorization_code")

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), err
}