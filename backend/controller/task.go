package controller

import (
	restApiModel "golang-nextjs-app/restapi/models"
	"golang-nextjs-app/restapi/operations"
	"golang-nextjs-app/usecase"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type TaskController struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (c *TaskController) CreateTask(params operations.PostTaskParams) middleware.Responder {

	var (
		ok = operations.NewPostTaskCreated()
	)

	ctx := params.HTTPRequest.Context()
	token := params.HTTPRequest.Header.Get("Authorization")

	if _, err := c.taskUsecase.CreateTask(ctx, params.Title, params.Content, token); err != nil {
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
