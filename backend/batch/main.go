package main

import (
	"context"
	"fmt"

	"github.com/Takenari-Yamamoto/golang-nextjs-app/db"
	"github.com/Takenari-Yamamoto/golang-nextjs-app/db/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func insertTask(count int) {
	db, err := db.NewDb()
	if err != nil {
		fmt.Println("DB エラー", err)
		return
	}
	defer db.Close()
	task := &models.Task{
		Title:       fmt.Sprintf("Dummy New Task %d", count),
		Description: null.StringFrom("This is a new task"),
	}
	err = task.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		fmt.Println("挿入失敗", err)
		return
	}

	fmt.Println("User inserted successfully.")
}

func main() {
	for i := 0; i < 100; i++ {
		insertTask(i)
	}
}
