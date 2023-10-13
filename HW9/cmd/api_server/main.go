package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	authenticated    map[string]bool = make(map[string]bool)
	Tasks_data       Tasks
	SESSION_COOKIE                      = "session"
	CONTEXT_AUTH_KEY                    = "authenticated"
	registered       map[login]password = make(map[login]password)
	logined_users    []login
)

func tast1() {
	filepath := "assets/tasks.json"
	Tasks_data = Read_Tasks(filepath)
	Read_Users("assets/users.json")
	r := mux.NewRouter()
	r.HandleFunc("/", Root)
	r.HandleFunc("/students", Students_path)
	r.HandleFunc("/tasks", Tasks_path)
	r.HandleFunc("/status", Status)
	r.HandleFunc("/task/bydate/{date}", GetTaskByDate)
	r.HandleFunc("/login", ApplyMiddleware(signIn, Authenticate)).Methods("POST")
	r.HandleFunc("/student/{id}", ApplyMiddleware(StundentByID, Authorize())).Methods("GET")
	http.ListenAndServe("localhost:8080", r)

}

func main() {
	tast1()
}
