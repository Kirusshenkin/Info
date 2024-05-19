package main

import (
	"cryptoApi/internal/database"
	"cryptoApi/pkg"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// connecting monoDB it is database.go
	database.Connect()

	// starting bot it is bot.go
	bot.Start()

	// logger info and error it is server.go
	log.Println("Server started")
}
