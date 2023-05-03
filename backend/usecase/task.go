package usecase

import (
	"context"
	"golang-nextjs-app/domain"
	"golang-nextjs-app/gateway/postgres"

	"firebase.google.com/go/v4/auth"
)

type TaskUsecase struct {
	repo       postgres.ITaskRepository
	authClient *auth.Client
}

func NewTaskUsecase(
	repo postgres.ITaskRepository,
	authClient *auth.Client,
) *TaskUsecase {
	return &TaskUsecase{
		repo:       repo,
		authClient: authClient,
	}
}

func (t *TaskUsecase) CreateTask(ctx context.Context, title, content, token string) (*string, error) {

	auth, err := t.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	res, err := t.repo.CreateTask(ctx, title, content, auth.UID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *TaskUsecase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := t.repo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskUsecase) GetTaskByID(ctx context.Context, id string) (*domain.Task, error) {
	task, err := t.repo.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskUsecase) UpdateTask(ctx context.Context, id, title, content string) error {
	err := t.repo.UpdateTask(ctx, id, title, content)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskUsecase) DeleteTask(ctx context.Context, id string) error {
	err := t.repo.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
