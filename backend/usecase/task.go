package usecase

import (
	"context"
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
