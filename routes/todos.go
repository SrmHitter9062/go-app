package routes

import (
	"github.com/SrmHitter9062/go-todo-app/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
	//"strconv"
)

func AddTodosRoutes(baseRouter gin.IRouter){
	v1 := baseRouter.Group("/api/v1/todo")
	{
		v1.GET("/",fetchAllTodo)
		v1.GET("/:id",fetchSingleTodo)
		v1.POST("/",createTodo)
		v1.PUT("/:id",updateTodo)
		v1.DELETE("/",deleteTodo)
	}
	return
}

func createTodo(c *gin.Context){
	var isCreated = controllers.CreateTodo(c)
	if isCreated {
		c.JSON(http.StatusCreated,gin.H{"status" : http.StatusCreated,"message": "todo item is created successfully","resourceId":isCreated})
	}else {
		c.JSON(http.StatusCreated,gin.H{"status" : http.StatusCreated,"message": "todo item is  not created","resourceId":isCreated})
	}
}

func fetchAllTodo(c *gin.Context){
	var data = controllers.FetchAllTodo(c)
	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":data})
}

func fetchSingleTodo(c *gin.Context){
	//var todo todoModel
	//todoId := c.Param("id")
	//db.First(&todo,todoId) // finds the first record where
	//if todo.ID == 0{
	//	c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
	//	return
	//}
	//completed := false
	//if todo.Completed == 1{
	//	completed = true
	//}else{
	//	completed = false
	//}
	//_todo := transformedTodo{ID:todo.ID,Title:todo.Title,Completed:completed}
	//c.JSON(http.StatusOK,gin.H{"status":http.StatusOK, "data":_todo})
	//return
}

func deleteTodo(c *gin.Context){
	//var todo todoModel
	//todoId := c.Param("id")
	//db.First(&todo,todoId)
	//if todo.ID == 0 {
	//	c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
	//	return
	//}
	//db.Delete(&todo)
	//c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"todo is deleted"})
}

func updateTodo(c *gin.Context){
	//var todo todoModel
	//todoId := c.Param("id")
	//db.First(&todo,todoId)
	//if todo.ID == 0 {
	//	c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
	//	return
	//}
	//db.Model(&todo).Update("title",c.PostForm("title"))
	//completed, _ := strconv.Atoi(c.PostForm("completed"))
	//db.Model(&todo).Update("completed", completed)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})

}
