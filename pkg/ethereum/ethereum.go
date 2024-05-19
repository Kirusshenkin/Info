package ethereum

import (
	"cryptoApi/pkg/bybit"
	"fmt"
	"log"
)

func GetPrice() {
	price, err := bybit.GetETHPrice()
	if err != nil {
		log.Printf("Error getting Ethereum price: %v", err)
		return
	}
	fmt.Printf("Ethereum Price: %s\n", price)
}
