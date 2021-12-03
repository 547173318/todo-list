package router

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() (r *gin.Engine){
	r = gin.Default()
	userGroup := r.Group("/v1")
	{
		userGroup.GET("/todo", controller.GetTodoList)
		userGroup.POST("/todo", controller.CreatTodo)
		userGroup.PUT("/todo/:id", controller.UpdateTodo)
		userGroup.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}
