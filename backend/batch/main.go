package main

import (
	"context"
	"fmt"

	"github.com/Takenari-Yamamoto/golang-nextjs-app/db"
	"github.com/Takenari-Yamamoto/golang-nextjs-app/repository"
)

func main() {
	db, err := db.NewDb()
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
