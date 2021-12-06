package client

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ApiBaseURL             string = "https://semsportal.com/api"
	LoginURL               string = "/v1/Common/CrossLogin"
	ExportPowerStationURL  string = "/v1/PowerStation/ExportPowerstationPac"
	GetStationPowerDataURL string = "/v1/ReportData/GetStationPowerDataFilePath"
	GetCurrentReadings     string = "/v1/PowerStation/GetMonitorDetailByPowerstationId"
)

func MakeRequest(url string, method string, headers http.Header, requestBody io.Reader) string {
	client := &http.Client{}
	req, _ := http.NewRequest(method, ApiBaseURL+url, requestBody)
	req.Header = headers

	res, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	//Read the data_transfer_object body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
