package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Ticker struct {
	Symbol    string `json:"symbol"`
	LastPrice string `json:"lastPrice"`
}

type TickerResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []Ticker `json:"list"`
	} `json:"result"`
}

func GetCryptoPrices(symbols []string) (map[string]string, error) {
	apiKey := os.Getenv("BYBIT_API_TOKEN")
	secret := os.Getenv("BYBIT_API_SECRET")
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	prices := make(map[string]string)

	// symbolsParam := strings.Join(symbols, ",")
	endpoint := fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=spot")
	sign := generateSignature(apiKey, secret, timestamp)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-BAPI-API-KEY", apiKey)
	req.Header.Set("X-BAPI-TIMESTAMP", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-BAPI-SIGN", sign)

	// Log the request details
	log.Printf("Sending request to Bybit API: %s\n", endpoint)
	log.Printf("Headers: X-BAPI-API-KEY: %s, X-BAPI-TIMESTAMP: %d, X-BAPI-SIGN: %s\n", apiKey, timestamp, sign)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response code: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Log the full response body
	log.Printf("Full response from Bybit API: %s\n", string(body))

	var tickerResponse TickerResponse
	if err := json.Unmarshal(body, &tickerResponse); err != nil {
		// Log the error and the body
		log.Printf("Error unmarshaling JSON: %v\n", err)
		log.Printf("Response body: %s\n", string(body))
		return nil, err
	}

	if tickerResponse.RetCode != 0 {
		return nil, fmt.Errorf("API error: %s", tickerResponse.RetMsg)
	}

	for _, item := range tickerResponse.Result.List {
		prices[item.Symbol] = item.LastPrice
	}

	return prices, nil
}

func generateSignature(apiKey, secret string, timestamp int64) string {
	data := fmt.Sprintf("%d%s", timestamp, apiKey)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
