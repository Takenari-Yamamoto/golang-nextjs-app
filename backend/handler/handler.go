package handler

import (
	"golang-nextjs-app/controller"
	"golang-nextjs-app/restapi/operations"
)

func HandleRestApi(api *operations.GolangNextjsAPI) {
	// db, err := database.NewDb()
	// if err != nil {
	// 	fmt.Println("failed to connect database:", err)
	// 	return
	// }
	// DB repository
	// // userRepo := postgres.NewUserRepository(db)
	// taskRepo := postgres.NewTaskRepository(db)

	// // usecase
	// taskUsecase := usecase.NewTaskUseCase(db, taskRepo)

	taskController := controller.NewTaskController()

	api.GetTasksHandler = operations.GetTasksHandlerFunc(taskController.GetAllTasks)
	// api.GetTasksIDHandler = operations.GetTasksIDHandlerFunc(taskController.GetTaskById)
}
