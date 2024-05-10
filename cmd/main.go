package main

import (
	"cryptoApi/internal/database"
	"cryptoApi/internal/handler"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
    database.Connect()

    http.HandleFunc("/", handler.HomeHandler)
    http.HandleFunc("/prices", handler.PricesHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))

		cwd, _ := os.Getwd()
		files, _ := ioutil.ReadDir("./")
		log.Printf("Current working directory: %s", cwd)
		for _, f := range files {
				log.Println(f.Name())
		}
}
