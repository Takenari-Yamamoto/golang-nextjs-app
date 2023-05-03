package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang-nextjs-app/database/models"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"golang-nextjs-app/domain"
)

type ITaskRepository interface {
	CreateTask(ctx context.Context, title string, content string, uid string) (*string, error)
	GetAllTasks(ctx context.Context) ([]*domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*domain.Task, error)
	UpdateTask(ctx context.Context, id string, title string, content string) error
	DeleteTask(ctx context.Context, id string) error
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

func (t *TaskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := models.Tasks().All(ctx, t.db)
	if err != nil {
		return nil, err
	}
	var domainTasks []*domain.Task

	for _, task := range tasks {
		domainTasks = append(domainTasks, convToDomainTask(task))
	}
	return domainTasks, nil
}

func (t *TaskRepository) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	task, err := models.FindTask(ctx, t.db, id)
	if err != nil {
		return nil, err
	}
	return convToDomainTask(task), nil
}

func (t *TaskRepository) UpdateTask(ctx context.Context, id string, title string, content string) error {
	task, err := models.FindTask(ctx, t.db, id)
	if err != nil {
		return err
	}
	task.Title = title
	task.Content = content
	_, err = task.Update(ctx, t.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	task, err := models.FindTask(ctx, t.db, id)
	if err != nil {
		return err
	}
	_, err = task.Delete(ctx, t.db)
	if err != nil {
		return err
	}
	return nil
}

func convToDomainTask(task *models.Task) *domain.Task {
	return &domain.Task{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Done:      task.Done,
		CreatedAt: task.Createdat,
		CreatedBy: task.Createdby,
	}
}
