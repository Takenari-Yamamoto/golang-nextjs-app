package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WebhookData struct {
	// 受け取りたいデータの構造を定義します。
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/webhook-endpoint", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}

		var data WebhookData
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error unmarshaling JSON", http.StatusInternalServerError)
			return
		}

		// Bodyを解析
		log.Default().Printf("Received message: %s", data.Message)

	})

	log.Default().Println("Server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
