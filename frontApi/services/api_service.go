package services

import (
	"encoding/json"
	"github.com/HETIC-MT-P2021/PROJECT_FINAL_GROUP03/frontApi/env"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func PerformApiRequest(route string, method string, data interface{}) (*http.Response, error) {
	var response *http.Response
	client := &http.Client{}
	baseUrl := env.GetVariable("DOMAIN_API_URL")
	requestBody, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	url := baseUrl + string(route)
	r, err := http.NewRequest(method, url, strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}

	r.Header.Set("content-type", "application/json")
	response, err = client.Do(r)

	return response, err
}

func DeserializeResponseFromObject(data io.ReadCloser, object interface{}) error {
	deserializedData, err := ioutil.ReadAll(data)
	if err != nil {
		log.Error("could ot read response for deserialization")
		return err
	}

	err = json.Unmarshal(deserializedData, &object)
	if err != nil {
		log.Error("could not unmarshal response data to object")
		return err
	}
	return nil
}
