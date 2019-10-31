package controllers

import (
	"github.com/SrmHitter9062/go-todo-app/helpers"
	"github.com/SrmHitter9062/go-todo-app/models/todo"
	"github.com/gin-gonic/gin"
)


func FetchAllTodo(c *gin.Context) interface{}{
	offset := c.Param("offset")
	if offset == "" {
		offset = "10"
	}
	var todoModel = todo.Todo{}
	result := todoModel.FetchAllTodoList(offset)
	return result

}

func CreateTodo(c *gin.Context)bool{
	var paramInstance = helpers.Params{} // params instance
	var todoModel = todo.Todo{} // todoModel instance
	var paramObj = paramInstance.GetParamObj(c)
	isCreated := todoModel.CreateTodo(paramObj)
	return isCreated
}
