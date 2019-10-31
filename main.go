package main

import (
	"github.com/SrmHitter9062/go-todo-app/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
func main(){
	router := gin.Default()
	routes.AddTodosRoutes(router)
	routes.AddUsersRoutes(router)
	router.Run()
}


func init(){
	//var err error
	//_,err = sqlx.Connect("mysql","root:@/todo?charset=utf8&parseTime=True&loc=Local")
	////config.dbInstance = db
	//if err != nil {
	//	panic("failed to connect to db")
	//}else{
	//	fmt.Println("connected to db")
	//}
}

//type (
//	todoModel struct {
//		gorm.Model
//		Title  string `json:"title"`
//		Completed int `json:completed`
//	}
//	transformedTodo struct{
//		ID uint `json:"id"`
//		Title string `json:"title"`
//		Completed bool `json:"completed"`
//	}
//)

//
//func createTodo(c *gin.Context){
//	completed, _ := strconv.Atoi(c.PostForm("completed"))
//	todo := todoModel{Title:c.PostForm("title"),Completed:completed}
//	db.Save(&todo)
//	c.JSON(http.StatusCreated,gin.H{"status" : http.StatusCreated,"message": "todo item is created successfully","resourceId":todo.ID})
//}
//
//func fetchAllTodo(c *gin.Context){
//	var todos []todoModel
//	var _todos []transformedTodo
//	db.Find(&todos)
//	if len(todos) <= 0 {
//		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todos not found"})
//		return
//	}
//	// make good response
//	for _,item := range todos {
//		completed := false
//		if item.Completed == 1 {
//			completed = true
//		} else{
//			completed = false
//		}
//		_todos = append(_todos,transformedTodo{ID:item.ID,Title:item.Title,Completed:completed})
//	}
//	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"data":_todos})
//	return
//}
//
//func fetchSingleTodo(c *gin.Context){
//	var todo todoModel
//	todoId := c.Param("id")
//	db.First(&todo,todoId) // finds the first record where
//	if todo.ID == 0{
//		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
//		return
//	}
//	completed := false
//	if todo.Completed == 1{
//		completed = true
//	}else{
//		completed = false
//	}
//	_todo := transformedTodo{ID:todo.ID,Title:todo.Title,Completed:completed}
//	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK, "data":_todo})
//	return
//}
//
//func deleteTodo(c *gin.Context){
//	var todo todoModel
//	todoId := c.Param("id")
//	db.First(&todo,todoId)
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
//		return
//	}
//	db.Delete(&todo)
//	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"todo is deleted"})
//}
//
//func updateTodo(c *gin.Context){
//	var todo todoModel
//	todoId := c.Param("id")
//	db.First(&todo,todoId)
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"todo not found"})
//		return
//	}
//	db.Model(&todo).Update("title",c.PostForm("title"))
//	completed, _ := strconv.Atoi(c.PostForm("completed"))
//	db.Model(&todo).Update("completed", completed)
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
//
//}
