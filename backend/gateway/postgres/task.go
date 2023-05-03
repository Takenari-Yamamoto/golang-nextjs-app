package postgres

import (
	"context"
	"golang-nextjs-app/database"
	"time"

	"golang-nextjs-app/database/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ITaskRepository interface {
	CreateTask(ctx context.Context, title string, content string, uid string) (*string, error)
}

type TaskRepository struct{}

func NewTaskRepository() ITaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) CreateTask(ctx context.Context, title string, content string, uid string) (*string, error) {
	db, err := database.NewDb()
	if err != nil {
		return nil, err
	}

	uuidObj, _ := uuid.NewUUID()

	task := models.Task{
		ID:        uuidObj.String(),
		Title:     title,
		Content:   content,
		Done:      false,
		Createdby: uid,
		Createdat: time.Now(),
	}

	if err := task.Insert(ctx, db, boil.Infer()); err != nil {
		return nil, err
	}

	return &task.ID, nil

}
