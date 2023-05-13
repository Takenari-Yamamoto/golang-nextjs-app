package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	pid := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(ctx, pid)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}

	topic := client.Topic("my-topic")
	defer topic.Stop()

	msg := &pubsub.Message{
		Data: []byte("Hello, worldやで!"),
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		log.Fatalf("Could not publish message: %v", err)
	}

	log.Println("Published a message")
}
