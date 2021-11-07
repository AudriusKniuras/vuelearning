package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var exercises []Exercises
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./fitness.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/allhistory", displayAllHistory).Methods("GET", "OPTIONS")
	// Ordering important, must come before the other /exhistory
	myRouter.HandleFunc("/exhistory/create", createExHistory).Methods("POST")
	myRouter.HandleFunc("/exhistory/updateweight", updateExWeight).Methods("POST")
	myRouter.HandleFunc("/exhistory/updatereps", updateExReps).Methods("POST")
	myRouter.HandleFunc("/exhistory/{user}", returnUserHistory)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func displayAllHistory(w http.ResponseWriter, r *http.Request) {
	exercises = getAllHistory(db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(exercises)
}

func returnUserHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["user"]

	for _, ex := range exercises {
		if ex.User == key {
			json.NewEncoder(w).Encode(ex)
		}
	}
}

func createExHistory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var ex Exercises
	json.Unmarshal(reqBody, &ex)

	insertIntoHistory(db, ex)
	displayAllHistory(w, r)
}

func updateExReps(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ex Exercises
	json.Unmarshal(reqBody, &ex)

	updateRepsIntoHistory(db, ex)
	displayAllHistory(w, r)
}

func updateExWeight(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ex Exercises
	json.Unmarshal(reqBody, &ex)

	updateWeightintoHistory(db, ex)
	displayAllHistory(w, r)
}
