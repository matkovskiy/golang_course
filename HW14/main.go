package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ViewData struct {
	Title   string
	Message string
	Time    string
}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/", rootHandler).Methods("GET")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	// Start server
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

// controllers
func rootHandler(w http.ResponseWriter, r *http.Request) {

	message := "Welcome to the Home Page!"
	fmt.Println("Got message")
	renderTemplate(w, "html-0001.template", message) // View
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// It's a controller
	message := "Welcome to the About Page!"
	fmt.Println("Got message - about")
	renderAbout(w, "html-0002.template", message) // View
}

func renderTemplate(w http.ResponseWriter, template_name, message string) {
	tmpl, err := template.New(template_name).ParseFiles("assets/" + template_name)
	if err != nil {
		panic(err)
	}
	time_now := time.Now()

	data := ViewData{
		Title:   "Welcome message",
		Message: message,
		Time:    time_now.String(),
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

// View
func renderAbout(w http.ResponseWriter, template_name, message string) {
	tmpl, err := template.New(template_name).ParseFiles("assets/" + template_name)
	if err != nil {
		panic(err)
	}
	time_now := time.Now()

	data := ViewData{
		Message: "It's about page",
		Time:    time_now.String(),
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
