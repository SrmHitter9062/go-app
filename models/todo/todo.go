package todo

import (
	"fmt"
	"github.com/SrmHitter9062/go-todo-app/db"
	"github.com/SrmHitter9062/go-todo-app/helpers"
)

type Todo struct {
	Id  int `json:todo_id`
	Title  string `json:"title"`
	IsCompleted int `json:is_completed`
	Status  string `json:status`
}
// FetchAllTodoList function on Todo
func (mytodo Todo)FetchAllTodoList(offset string ) interface{}{
	var dbInstance = db.GetInstance()
	res,err := dbInstance.ReadInstance.Query("SELECT id,title,is_completed ,status FROM todos LIMIT ?",offset)
	if err != nil {
		panic(err.Error())
	}
	var finalResult = []Todo{}
	for res.Next() {
		var title string
		var is_completed,id int
		var status string
		err = res.Scan(&id,&title, &is_completed, &status)
		finalResult = append(finalResult, Todo{Title:title,Status:status,IsCompleted:is_completed})
		//fmt.Println(res)
	}
	return finalResult
}
//creatTodo function on struct Todo
func (mytodo Todo)CreateTodo(data helpers.Params ) bool{
	fmt.Println(data.Title)
	var dbInstance = db.GetInstance()
	res,err := dbInstance.ReadInstance.Query("insert into todos(title,is_completed,status) values(?,?,?)",data.Title,data.IsCompleted,data.Status)
	if(err != nil){
		//panic(err.Error())
		return false
	}
	fmt.Println(res)
	return true
}
