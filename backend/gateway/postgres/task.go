package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang-nextjs-app/database/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ITaskRepository interface {
	CreateTask(ctx context.Context, title string, content string, uid string) (*string, error)
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(
	db *sql.DB,
) ITaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) CreateTask(ctx context.Context, title string, content string, uid string) (*string, error) {

	uuidObj, _ := uuid.NewUUID()

	task := models.Task{
		ID:        uuidObj.String(),
		Title:     title,
		Content:   content,
		Done:      false,
		Createdby: uid,
		Createdat: time.Now(),
	}

	fmt.Println(task)

	if err := task.Insert(ctx, t.db, boil.Infer()); err != nil {
		return nil, err
	}
	return &task.ID, nil
}
