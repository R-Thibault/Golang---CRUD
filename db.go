package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=thibault password=superPassword dbname=bookstore sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database: ", err)
	}
	fmt.Println("Connected to the database successfully!")
}