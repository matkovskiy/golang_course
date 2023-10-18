package main

import (
	"encoding/json"
	"fmt"
	"hw10/internal/structs"
	"hw10/pkg/external_api"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WeaterCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusBadRequest)
	response := external_api.GetWeatherByCity(vars["city"], WeatherApiKey)
	w.Write([]byte(response))
}

func Root(w http.ResponseWriter, r *http.Request) {
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(map[string]any{"What's up": "bor"})
	w.Write(response)
}

func TranslateMe(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var text structs.Text
	err = json.Unmarshal(body, &text)
	fmt.Printf("Source language = %s\n", text.Source_language)
	fmt.Printf("Target language = %s\n", text.Target_language)
	fmt.Printf("Text = %s\n", text.Text)
	output := external_api.Translate(text.Source_language, text.Target_language, text.Text)
	response, _ := json.Marshal(map[string]any{"Output": output})
	w.Write(response)
}
