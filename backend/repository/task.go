package repository

import (
	"context"
	"database/sql"

	"github.com/Takenari-Yamamoto/golang-nextjs-app/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoRepoMethods interface {
	GetAll(ctx context.Context) ([]*models.Task, error)
	Create(ctx context.Context, title, description string) error
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) TodoRepoMethods {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) GetAll(ctx context.Context) ([]*models.Task, error) {
	tasks, err := models.Tasks().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TodoRepo) Create(ctx context.Context, title, description string) error {
	task := &models.Task{
		Title: title,
	}
	err := task.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
