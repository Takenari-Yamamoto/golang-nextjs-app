package main

import (
	"context"
	"fmt"

	"github.com/Takenari-Yamamoto/golang-nextjs-app/database"
	"github.com/Takenari-Yamamoto/golang-nextjs-app/database/repository"
)

func main() {
	db, err := database.NewDb()
	if err != nil {
		return
	}

	repo := repository.NewTodoRepo(db)
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		if err := repo.Create(ctx, fmt.Sprintf("Dummy New Task %d", i), "This is a new task"); err != nil {
			fmt.Println("err YADE")
		}
	}
}
