package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(logined_users)
	w.Write(response)
}
