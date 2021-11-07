package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func getAllHistory(db *sql.DB) []Exercises {
	var exercises []Exercises
	rows, err := db.Query("SELECT * FROM exhistory")
	if err != nil {
		log.Fatal(err)
	}
	var ex Exercises
	for rows.Next() {
		err = rows.Scan(&ex.Id, &ex.User, &ex.Date, &ex.Program, &ex.Day, &ex.Exercise, &ex.Sets, &ex.Reps, &ex.Weight)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal([]byte(ex.Reps), &ex.RepsArr)
		exercises = append(exercises, ex)
	}
	return exercises
}

func insertIntoHistory(db *sql.DB, ex Exercises) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	insert, err := tx.Prepare("INSERT INTO exhistory (email,date,program,day,exercise,sets,reps,weight) VALUES (?,?,?,?,?,?,?,?)")

	fmt.Println(insert)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

	_, err = insert.Exec(ex.User, ex.Date, ex.Program, ex.Day, ex.Exercise, ex.Sets, ex.Reps, ex.Weight)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func updateRepsIntoHistory(db *sql.DB, ex Exercises) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	update, err := tx.Prepare("UPDATE exhistory SET reps = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()

	_, err = update.Exec(ex.Reps, ex.Id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func updateWeightintoHistory(db *sql.DB, ex Exercises) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	update, err := tx.Prepare("UPDATE exhistory SET weight = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()

	_, err = update.Exec(ex.Weight, ex.Id)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
