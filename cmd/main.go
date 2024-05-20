package main

import (
	"cryptoApi/internal/database"
	bot "cryptoApi/pkg"
	"cryptoApi/pkg/bybit"
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// connecting to MongoDB in database.go
	database.Connect()

	// check for command line arguments
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "crypto":
			getCryptoPrices()
		default:
			fmt.Println("Unknown command")
		}
	} else {
		// starting bot in bot.go
		bot.Start()

		// logger info and error in server.go
		log.Println("Server started")
	}
}

func getCryptoPrices() {
	symbols := []string{"BTCUSDT", "ETHUSDT", "SOLUSDT", "ATOMUSDT", "ADAUSDT"}
	prices, err := bybit.GetCryptoPrices(symbols)
	if err != nil {
		fmt.Printf("Error getting crypto prices from Bybit: %v\n", err)
		return
	}
	for symbol, price := range prices {
		fmt.Printf("Bybit: %s: %s$\n", symbol, price)
	}
}
