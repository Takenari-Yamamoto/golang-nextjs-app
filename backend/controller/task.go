package controller

import (
	"fmt"
	"golang-nextjs-app/domain"
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

	var restApiTasks []*restApiModel.Task
	tasks, err := c.taskUsecase.GetAllTasks(params.HTTPRequest.Context())
	if err != nil {
		fmt.Println("タスクの取得に失敗しました ---->>>>", err)
		return operations.NewGetTasksInternalServerError()
	}
	for _, task := range tasks {
		restApiTasks = append(restApiTasks, convertTaskToRestApiTask(task))
	}

	return ok.WithPayload(restApiTasks)
}

// domain.Taskをrestapi.Taskに変換する
func convertTaskToRestApiTask(task *domain.Task) *restApiModel.Task {
	return &restApiModel.Task{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Done:      task.Done,
		CreatedAt: strfmt.DateTime(task.CreatedAt),
		CreatedBy: task.CreatedBy,
	}
}
