package main

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	pid := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(ctx, pid)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	sub := client.Subscription("my-sub")
	cctx, cancel := context.WithCancel(ctx)

	var mu sync.Mutex
	sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()

		log.Printf("Received message: %s", string(msg.Data))

		msg.Ack()

		cancel()
	})
}
