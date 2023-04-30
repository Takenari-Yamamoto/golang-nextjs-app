package controller

import (
	restApiModel "golang-nextjs-app/restapi/models"
	"golang-nextjs-app/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (c *TaskController) GetAllTasks(params operations.GetTasksParams) middleware.Responder {
	var (
		ok = operations.NewGetTasksOK()
	)

	var tasks []*restApiModel.Task
	for i := 0; i < 30; i++ {
		u, _ := uuid.NewRandom()
		tasks = append(tasks, &restApiModel.Task{
			ID:        u.String(),
			Title:     "dummy title",
			Content:   "dummy content",
			Done:      false,
			CreatedAt: strfmt.DateTime{},
			CreatedBy: "TAKENARI",
		})
	}

	return ok.WithPayload(tasks)
}
