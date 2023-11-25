package main

import (
	"SKLAD/pkg/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

// controllers
func time_now_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Got message")
	w.Write([]byte(db.Db_timeNow()))

}
func db_ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(db.Db_ping_log()))

}

func all_products(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(db.GetAllProducts()))
}

func get_product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := db.Cache_Get("product:" + vars["product_id"])

	if err == redis.Nil {
		fmt.Println("key does not exist in cache")
		result = db.GetProductByIP(vars["product_id"])
		db.Cache_Set("product:"+vars["product_id"], result)
	} else {
		fmt.Println("returning from cache")
	}

	w.Write([]byte(result))

}

func main() {
	db.Db_init()

	router := mux.NewRouter()
	// Routes
	router.HandleFunc("/", time_now_handler).Methods("GET")
	router.HandleFunc("/ping", db_ping).Methods("GET")
	router.HandleFunc("/products/all", all_products).Methods("GET")
	router.HandleFunc("/product/{product_id}", get_product).Methods("GET")
	http.ListenAndServe("localhost:8080", router)

}
