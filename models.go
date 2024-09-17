package main

type Config struct {
	Database struct {
			User     string `json:"user"`
			Password string `json:"password"`
			Name     string `json:"name"`
			Host     string `json:"host"`
			Port     int    `json:"port"`
	} `json:"database"`
	Server struct {
			Port int `json:"port"`
	} `json:"server"`
}

type Book struct {
	ID int `json:"id"`
	Title	string	`json:"title"`
	Author string `json:"author"`
}

