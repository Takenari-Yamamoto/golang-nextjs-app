package controller

import (
	"fmt"
	restApiModel "golang-nextjs-app/restapi/models"
	"golang-nextjs-app/restapi/operations"
	"golang-nextjs-app/usecase"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type TaskController struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskController(
	taskUsecase *usecase.TaskUsecase,
) *TaskController {
	return &TaskController{
		taskUsecase: taskUsecase,
	}
}

func (c *TaskController) CreateTask(params operations.PostTaskParams) middleware.Responder {
	var (
		ok = operations.NewPostTaskCreated()
	)

	ctx := params.HTTPRequest.Context()
	bt := params.HTTPRequest.Header.Get("Authorization")
	idToken := strings.TrimPrefix(bt, "Bearer ")
	b := params.Body

	fmt.Println("token です", idToken)

	if _, err := c.taskUsecase.CreateTask(ctx, *b.Title, *b.Content, idToken); err != nil {
		fmt.Println("タスクの作成に失敗しました ---->>>>", err)
		return operations.NewPostTaskInternalServerError()
	}

	return ok

}

func (c *TaskController) GetAllTasks(params operations.GetTasksParams) middleware.Responder {
	var (
		ok = operations.NewGetTasksOK()
	)

	var tasks []*restApiModel.Task
	for i := 0; i < 30; i++ {
		tasks = append(tasks, &restApiModel.Task{
			ID:        "id" + string(i),
			Title:     "title",
			Content:   "content",
			Done:      false,
			CreatedAt: strfmt.DateTime{},
			CreatedBy: "TAKENARI",
		})
	}

	return ok.WithPayload(tasks)
}
