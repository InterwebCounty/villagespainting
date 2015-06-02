package todoroutes

import (
	"github.com/interwebcounty/villagespainting/server/api/todo/controller"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	router.GET("/api/todos", todocontroller.GetAll)
	router.POST("/api/todos", todocontroller.NewTodo)
	router.DELETE("/api/todos/:id", todocontroller.RemoveTodo)
}
