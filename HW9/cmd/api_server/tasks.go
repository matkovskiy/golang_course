package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Tasks_path(w http.ResponseWriter, r *http.Request) {
	filepath := "assets/tasks.json"
	Tasks := Read_Tasks(filepath)

	fmt.Println(Tasks)
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(Tasks)
	w.Write(response)
}

func GetTaskByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filepath := "assets/tasks.json"
	Tasks := Read_Tasks(filepath)

	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(vars["date"])
	for i := range Tasks.Task {
		if Tasks.Task[i].Date == vars["date"] {
			response, _ := json.Marshal(Tasks.Task[i])
			w.Write(response)
			break
		}
	}
}
