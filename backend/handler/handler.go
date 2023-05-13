package handler

import (
	"context"
	"golang-nextjs-app/controller"
	"golang-nextjs-app/database"
	"golang-nextjs-app/gateway/firebase"
	"golang-nextjs-app/gateway/postgres"
	"golang-nextjs-app/restapi/operations"
	"golang-nextjs-app/usecase"
	"log"
)

func HandleRestApi(api *operations.GolangNextjsAPI) {

	ctx := context.Background()
	db, err := database.NewDb()
	if err != nil {
		return
	}
	app := firebase.InitFirebase()
	auth, err := app.Auth(ctx)
	if err != nil {
		return
	}

	// データベースに接続できるか確認
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	taskRepo := postgres.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, auth)
	taskController := controller.NewTaskController(
		taskUsecase,
	)

	api.GetTasksHandler = operations.GetTasksHandlerFunc(taskController.GetAllTasks)
	api.PostTaskHandler = operations.PostTaskHandlerFunc(taskController.CreateTask)
	api.GetTasksIDHandler = operations.GetTasksIDHandlerFunc(taskController.GetTaskByID)
	api.DeleteTasksIDHandler = operations.DeleteTasksIDHandlerFunc(taskController.DeleteTask)
}
