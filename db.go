package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

var db *sql.DB

func initDB() {
	// Charge config from config.json
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal("Erreur lors du chargement de la configuration: ", err)
	}

	// Construct connexion with config.json value
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.Database.User, config.Database.Password, config.Database.Name, config.Database.Host, config.Database.Port)

	// Open BDD connexion
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Test BDD connexion
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database: ", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Table "books" creation
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100),
	author VARCHAR(100)
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
	fmt.Println("Table 'books' initialized succesfully!")
}
