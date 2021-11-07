package main

import (
	_ "github.com/mattn/go-sqlite3"
)

type Exercises struct {
	Id       int     `json:"id"`
	User     string  `json:"user"`
	Date     string  `json:"date"`
	Program  string  `json:"program"`
	Day      int     `json:"day"`
	Exercise string  `json:"exercise"`
	Sets     int     `json:"sets"`
	Reps     string  `json:"reps"`
	RepsArr  []int   `json:"repsArr"`
	Weight   float32 `json:"weight"`
}
