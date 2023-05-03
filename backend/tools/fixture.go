package tools

import (
	"context"
	"fmt"
	"golang-nextjs-app/gateway/firebase"
	"log"
	"sync"
)

func CreateDummyAuthUser(ctx context.Context) {
	client := firebase.NewFirebase()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			email := fmt.Sprintf("dummy%d@dummy.com", i)
			res, err := client.CreateUser(ctx, email, "password")
			if err != nil {
				log.Printf("failed to create user %s: %v", email, err)
				return
			}
			log.Printf("created user %s", res.Email)
		}(i)
	}
	wg.Wait()
}
