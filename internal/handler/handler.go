package handler

import (
	"cryptoInfo/internal/prices"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl.Execute(w, nil)
}

func PricesHandler(w http.ResponseWriter, r *http.Request) {
    pricesData, err := prices.GetCryptoPrices()
    if err != nil {
        http.Error(w, "Unable to fetch prices", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(pricesData)
}
