package commands

import (
	"bytes"
	"encoding/json"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/enum"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/bot/env"
	"net/http"
)

func PerformRequest(route enum.Routes, method string, data interface{}) error {
	client := &http.Client{}
	baseUrl := env.GetVariable("DOMAIN_API_URL")
	requestBody, err := json.Marshal(data)

	if err != nil {
		return err
	}

	url := baseUrl + string(route)
	r, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}
	_, err = client.Do(r)

	return err
}