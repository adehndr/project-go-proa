package app

import (
	"example.com/adehndr/project_go_proa/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(taskController controller.TaskController) *httprouter.Router{
	router := httprouter.New()

	router.GET("/tasks",taskController.FindAll)
	router.GET("/task/:taskid",taskController.FindById)
	router.POST("/tasks",taskController.Create)
	router.PUT("/task/:taskid",taskController.Update)
	router.DELETE("/task/:taskid",taskController.Delete)
	return router
}