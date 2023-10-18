package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Read_Users(filename string) {
	var Users Users
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Users)
	fmt.Println(Users.User[0].Name)
	for i := range Users.User {
		h := sha1.New()
		h.Write([]byte(Users.User[i].Password))

		pass := password(hex.EncodeToString(h.Sum(nil)))
		registered[login(Users.User[i].Name)] = pass
	}
	fmt.Println(registered)

	defer jsonFile.Close()
	// return
}

func Read_Student(filename string) (result Students) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	defer jsonFile.Close()
	return
}

func Read_Tasks(filename string) (result Tasks) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	defer jsonFile.Close()
	return
}
