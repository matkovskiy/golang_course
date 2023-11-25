package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Product struct {
	id                  int    `json:"id"`
	short_name          string `json:"short_name"`
	description         string `json:"description"`
	product_part_number string `json:"product_part_number"`
}

func Db_init() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "dbpass",
		Net:    "tcp",
		Addr:   "127.0.0.1:13306",
		DBName: "sklad",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

}

func Db_timeNow() string {
	type timeNow struct {
		time string
	}
	// Execute the query
	results, err := db.Query("SELECT NOW()")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var time_now timeNow
	for results.Next() {

		err = results.Scan(&time_now.time)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(time_now.time)
	return time_now.time
}

func Db_ping_log() string {
	sql := "INSERT INTO ping_log(ping_time) VALUES (NOW())"
	res, err := db.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)
	return strconv.FormatInt(lastId, 10)
}

func Get_products() []Product {

	sql := "select id, short_name, description, product_part_number from product;"
	res, err := db.Query(sql)

	if err != nil {
		log.Fatal(err)
	}
	var result []Product
	if res.Next() {

		var res_product Product
		err := res.Scan(&res_product.id, &res_product.short_name, &res_product.description, &res_product.product_part_number)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, res_product)

	}
	return result
}

func GetAllProducts() string {

	sql := "select id, short_name, description, product_part_number from product;"

	result, _ := getJSON(sql)

	return result
}

func getJSON(sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
