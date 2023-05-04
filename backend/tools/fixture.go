package tools

import (
	"context"
	"fmt"
	"golang-nextjs-app/database"
	"golang-nextjs-app/gateway/firebase"
	"golang-nextjs-app/gateway/postgres"
	"log"
	"strings"
	"sync"

	"firebase.google.com/go/v4/auth"
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

func createDummyTaskForUser(ctx context.Context, user *auth.ExportedUserRecord, taskRepo postgres.ITaskRepository, wg *sync.WaitGroup) {
	defer wg.Done()

	if !strings.Contains(user.Email, "dummy") {
		return
	}

	for i := 0; i < 10; i++ {
		_, err := taskRepo.CreateTask(ctx, "ダミータイトルです", "ダミー内容です", user.UID)
		if err != nil {
			log.Printf("failed to create task: %v", err)
			continue
		}
		log.Printf("created task: %s, %s", user.Email, user.UID)
	}
}

func CreateDummyTask(ctx context.Context) {
	client := firebase.NewFirebase()
	db, err := database.NewDb()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
		return
	}
	users, err := client.ListUsers(ctx)
	taskRepo := postgres.NewTaskRepository(db)
	if err != nil {
		return
	}
	if err != nil {
		log.Fatalf("failed to get users: %v", err)
		return
	}

	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go createDummyTaskForUser(ctx, user, taskRepo, &wg)
	}

	wg.Wait()
}
