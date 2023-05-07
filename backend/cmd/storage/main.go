package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("uploadHandler")
	// CORSヘッダーの追加
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Preflightリクエスト (OPTIONS) を処理
	if r.Method == http.MethodOptions {
		fmt.Println("Method: OPTIONS")
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		fmt.Println("Method: ", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		fmt.Println("FormFile: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("header: ", header)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("NewClient: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer client.Close()

	bucket := client.Bucket("demo-st")
	object := bucket.Object(uuid.New().String() + "-" + header.Filename)

	wc := object.NewWriter(ctx)
	wc.ContentType = header.Header.Get("Content-Type")
	if _, err = io.Copy(wc, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := wc.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", "demo-st", object.ObjectName())
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"url":"%s"}`, url)))
}

func main() {
	log.Println("server started on port 4200")
	http.HandleFunc("/api/upload", uploadHandler)
	if err := http.ListenAndServe(":4200", nil); err != nil {
		log.Fatal("failed to start http server", err)
		panic(err)
	}
}
