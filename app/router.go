package app

import (
	"example.com/adehndr/project_go_proa/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(taskController controller.TaskController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", taskController.Home)
	router.GET("/detail", taskController.Detail)
	router.POST("/detail", taskController.Detail)
	router.GET("/change/:taskid",taskController.Change)
	router.GET("/api/tasks", taskController.FindAll)
	router.GET("/api/task/:taskid", taskController.FindById)
	router.POST("/api/tasks", taskController.Create)
	router.PUT("/api/task/:taskid", taskController.Update)
	router.DELETE("/api/task/:taskid", taskController.Delete)
	return router
}
