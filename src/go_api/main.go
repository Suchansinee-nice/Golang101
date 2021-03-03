package main

import (
	"encoding/json"
	"fmt"
	"net/http" // import package net-http เข้ามา
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	fmt.Println("Create File Main.go Success")
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) { // (1)
	fmt.Fprint(w, "Welcome to the HomePage!") // (2)
}
func handleRequest() { // (3)
	http.HandleFunc("/", homePage) // (4)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":8080", nil) // (5)
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook) // (1)
}
