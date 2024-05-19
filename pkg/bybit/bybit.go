package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetETHPrice() (string, error) {
	apiKey := os.Getenv("BYBIT_API_TOKEN")
	secret := os.Getenv("BYBIT_API_SECRET")
	endpoint := "https://api.bybit.com/v5/market/ticker?symbol=ETHUSDT"

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	sign := generateSignature(apiKey, secret, timestamp)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-BAPI-API-KEY", apiKey)
	req.Header.Set("X-BAPI-TIMESTAMP", fmt.Sprintf("%d", timestamp))
	req.Header.Set("X-BAPI-SIGN", sign)

	// Log the request details
	log.Printf("Sending request to Bybit API: %s\n", endpoint)
	log.Printf("Headers: X-BAPI-API-KEY: %s, X-BAPI-TIMESTAMP: %d, X-BAPI-SIGN: %s\n", apiKey, timestamp, sign)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Log the response details
	log.Printf("Received response from Bybit API: %s\n", string(body))

	return string(body), nil
}

func generateSignature(apiKey, secret string, timestamp int64) string {
	data := fmt.Sprintf("%d%s", timestamp, apiKey)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
