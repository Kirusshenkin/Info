package prices

import (
	"encoding/json"
)

type CryptoPrice struct {
    Symbol string `json:"symbol"`
    Price  string `json:"price"`
}

func GetCryptoPrices() ([]byte, error) {
    // Здесь должен быть API-запрос к криптовалютной бирже.
    // Для примера, используем фиктивные данные.
    sampleData := CryptoPrice{
        Symbol: "BTCUSDT",
        Price: "48000.00",
    }
    jsonData, err := json.Marshal(sampleData)
    if err != nil {
        return nil, err
    }
    return jsonData, nil
}
