package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
)

var input Person

// create dummy data for insert to mongo
type Person struct {
	Name   string `bson:"name"`
	Number int    `bson:"number"`
}

func main() {

	http.HandleFunc("/add", add)
	http.ListenAndServe(":8080", nil)

}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p Person = Person{}
	decoder := json.NewDecoder(r.Body) //read from r

	err := decoder.Decode(&p) //Decode read JSON value from decoder store in p
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// connect to
	//ทำการเชื่อมต่อ mongo ด้วยคำสั่ง mgo.Dial โดยใส่ argument เป็น localhost:27017 โดย function นี้จะ return session กลับมา
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// select database dummy and collection test collection
	mongoConnection := session.DB("dummy").C("testCollection")
	// insert document
	err = mongoConnection.Insert(p)
	if err != nil {
		log.Fatal(err)
	}
	//close connection
	session.Close()

	defer r.Body.Close()

	//พ่น data กลับไปเมื่อ success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}
