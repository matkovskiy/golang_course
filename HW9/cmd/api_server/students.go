package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Students_path(w http.ResponseWriter, r *http.Request) {
	filepath := "assets/students.json"
	Students := Read_Student(filepath)
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(Students)
	w.Write(response)
}

func StundentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filepath := "assets/students.json"
	Students := Read_Student(filepath)

	w.WriteHeader(http.StatusBadRequest)
	fmt.Println(vars["id"])
	for i := range Students.Student {
		if Students.Student[i].ID == vars["id"] {
			response, _ := json.Marshal(Students.Student[i])
			w.Write(response)
			break
		}
	}
}
