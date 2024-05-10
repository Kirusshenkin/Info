package handler

import (
	"cryptoApi/internal/prices"
	"html/template"
	"log"
	"net/http"
)

func PricesHandler(w http.ResponseWriter, r *http.Request) {
    pricesData, err := prices.GetCryptoPrices()
    if err != nil {
        http.Error(w, "Unable to fetch prices", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(pricesData)
}

func init() {
    var err error
    tmpl, err = template.ParseFiles("templates/index.html")
    if err != nil {
        log.Fatalf("Failed to parse template: %v", err)
    }
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    priceData, err := prices.GetCryptoPrices()
    if err != nil {
        log.Printf("Error fetching prices: %v", err)
        http.Error(w, "Unable to fetch prices", http.StatusInternalServerError)
        return
    }

    data := struct {
        Price string
    }{
        Price: string(priceData),
    }

    if err := tmpl.Execute(w, data); err != nil {
        log.Printf("Error rendering template: %v", err)
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}
