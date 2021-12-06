package application

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"goodwe_monitor/application/data_transfer_object"
	"goodwe_monitor/domain/value_object"
	"goodwe_monitor/infrastructure/client"
	"log"
	"net/http"
)

var token value_object.Token
var account = GoodweAccount{}

type GoodweAccount struct {
	GoodweUser     string
	GoodwePassword string
}

func Login() bool {
	token = value_object.Token{Client: "ios", Version: "v2.0.4", Language: "en"}

	jsonToken, _ := json.Marshal(token)

	account.Read()

	postBody, _ := json.Marshal(map[string]string{
		"account": account.GoodweUser,
		"pwd":     account.GoodwePassword,
	})

	requestBody := bytes.NewBuffer(postBody)

	headers := http.Header{
		"User-Agent":      []string{"PVMaster/2.0.4 (iPhone; iOS 11.4.1; Scale/2.00)"},
		"Accept-Language": []string{"nl-BE;q=1"},
		"Content-Type":    []string{"application/json"},
		"Connect":         []string{"keep-alive"},
		"Token":           []string{string(jsonToken)},
	}

	response := client.MakeRequest(client.LoginURL, http.MethodPost, headers, requestBody)

	var loginResponse data_transfer_object.LoginResponse

	json.Unmarshal([]byte(response), &loginResponse)

	token = loginResponse.Data

	return true
}

func GetCurrentReadings() data_transfer_object.ReadingResponse {
	jsonToken, _ := json.Marshal(token)

	postBody, _ := json.Marshal(map[string]string{
		"powerStationId": "b7d59f76-3f42-469a-a8fd-33ff35eb6f85",
	})

	requestBody := bytes.NewBuffer(postBody)

	headers := http.Header{
		"User-Agent":      []string{"PVMaster/2.0.4 (iPhone; iOS 11.4.1; Scale/2.00)"},
		"Accept-Language": []string{"nl-BE;q=1"},
		"Content-Type":    []string{"application/json"},
		"Connect":         []string{"keep-alive"},
		"Token":           []string{string(jsonToken)},
	}

	response := client.MakeRequest(client.GetCurrentReadings, http.MethodPost, headers, requestBody)

	var readingResponse data_transfer_object.ReadingResponse

	json.Unmarshal([]byte(response), &readingResponse)

	return readingResponse
}

func (c *GoodweAccount) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
