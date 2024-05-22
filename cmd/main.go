package main

import (
	config "cryptoApi/internal/config"
	"cryptoApi/internal/database"
	bot "cryptoApi/pkg"
)

func main() {

	config.LoadEnv()

	database.Connect()
	bot.Start()
}
