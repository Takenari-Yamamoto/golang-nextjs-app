package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type WebhookData struct {
	// 送信したいデータの構造を定義します。
	Message string `json:"message"`
}

func main() {
	data := WebhookData{
		Message: "Hello, Webhook!",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Default().Fatalf("JSON marshaling failed: %s", err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/webhook-endpoint", body)
	if err != nil {
		log.Default().Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Default().Fatalf("Error sending request: %s", err)
	}

	defer resp.Body.Close()
}
