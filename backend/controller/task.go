package controller

import (
	"fmt"
	"golang-nextjs-app/domain"
	restApiModel "golang-nextjs-app/restapi/models"
	"golang-nextjs-app/restapi/operations"
	"golang-nextjs-app/usecase"
	"log"
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

	if _, err := c.taskUsecase.CreateTask(ctx, *b.Title, *b.Content, idToken); err != nil {
		fmt.Println("タスクの作成に失敗しました ---->>>>", err)
		return operations.NewPostTaskInternalServerError()
	}

	log.Println("タスクの作成に成功しました")

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
		return operations.NewGetTasksIDOK()
	}
	for _, task := range tasks {
		restApiTasks = append(restApiTasks, convertTaskToRestApiTask(task))
	}

	log.Println("タスクの取得に成功しました")

	return ok.WithPayload(restApiTasks)
}

func (c *TaskController) GetTaskByID(params operations.GetTasksIDParams) middleware.Responder {
	var (
		ok = operations.NewGetTasksIDOK()
	)

	ctx := params.HTTPRequest.Context()
	id := params.ID

	task, err := c.taskUsecase.GetTaskByID(ctx, id)
	if err != nil {
		fmt.Println("タスクの取得に失敗しました ---->>>>", err)
		return operations.NewGetTasksIDOK()
	}

	return ok.WithPayload(convertTaskToRestApiTask(task))
}

func (c *TaskController) DeleteTask(params operations.DeleteTasksIDParams) middleware.Responder {
	var (
		ok = operations.NewDeleteTasksIDOK()
	)

	ctx := params.HTTPRequest.Context()
	id := params.ID

	if err := c.taskUsecase.DeleteTask(ctx, id); err != nil {
		fmt.Println("タスクの削除に失敗しました ---->>>>", err)
		return operations.NewDeleteTasksIDOK()
	}

	return ok
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
