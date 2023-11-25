package main

import (
	"SKLAD/pkg/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// controllers
func timeNowHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Got message")
	w.Write([]byte(db.Db_timeNow()))

}
func dbping(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(db.Db_ping_log()))

}

func product(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(db.GetAllProducts()))

}
func main() {
	db.Db_init()

	router := mux.NewRouter()
	// Routes
	router.HandleFunc("/", timeNowHandler).Methods("GET")
	router.HandleFunc("/ping", dbping).Methods("GET")
	router.HandleFunc("/products", product).Methods("GET")

	http.Handle("/", router)
	http.Handle("/ping", router)
	http.Handle("/products", router)
	http.ListenAndServe(":8080", nil)

}
