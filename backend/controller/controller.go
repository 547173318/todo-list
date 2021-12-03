package controller

import (
	"backend/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodoList(c *gin.Context) {
	//获取所有的todo
	todoList,err := model.GetTodoList()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK,todoList)
	}

}

func CreatTodo(c *gin.Context) {
	//获取参数中的todo
	var todo model.Todo
	c.ShouldBind(&todo)

	//添加
	if err:= model.CreatTodo(&todo);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK,todo)
	}
}

func UpdateTodo(c *gin.Context) {
	//获取参数中的id
	idString,_:=c.Params.Get("id")
	idInt,_:= strconv.Atoi(idString)

	//查数据库得到原先的todo，这里是指针
	todo,_:= model.GetOne(uint(idInt))
	c.ShouldBind(todo)

	//更新
	if err := model.UpdateTodo(todo);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK,idInt)
	}
}

func DeleteTodo(c *gin.Context) {
	//获取参数中的id
	idString,_:=c.Params.Get("id")
	idInt,_:=strconv.Atoi(idString)

	//删除
	if err := model.DeleteTodo(uint(idInt));err !=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK,idInt)
	}
}