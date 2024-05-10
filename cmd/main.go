package main

import (
	"cryptoInfo/internal/database"
	"cryptoInfo/internal/handler"
	"log"
	"net/http"
)

func main() {
    database.Connect()

    http.HandleFunc("/", handler.HomeHandler)
    http.HandleFunc("/prices", handler.PricesHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
