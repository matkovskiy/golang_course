package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var WeatherApiKey string

func main() {
	WeatherApiKey = os.Getenv("WEATHER_API_KEY")
	if WeatherApiKey == "" {
		fmt.Println("Please set WEATHER_API_KEY")
		os.Exit(255)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/wather/{city}", WeaterCity)
	r.HandleFunc("/transalate", TranslateMe)

	http.ListenAndServe("localhost:8080", r)

}
