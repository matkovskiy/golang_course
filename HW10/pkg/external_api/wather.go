package external_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	api_url            = "https://api.weatherapi.com/v1/"
	curent_wather_path = "current.json"
	// api_key            = "XXXXXXXXXXXXXXXXXXXX"
	path_suffix = "&aqi=yes"
)

func GetWeatherByCity(city, api_key string) string {
	request_path := api_url + curent_wather_path + "?key=" + api_key + "&q=" + city + path_suffix
	response, err := http.Get(request_path)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}
