package helpers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Params struct {
	Title  string `json:"title"`
	IsCompleted int `json:is_completed`
	Status  string `json:status`
}


func (p Params)GetParamObj(c *gin.Context) Params {
	var title,status string
	var completed int
	title = c.PostForm("title")
	status = c.PostForm("status")
	completed,_ = strconv.Atoi(c.PostForm("completed"))
	var data = Params{Title:title, IsCompleted :completed,Status:status}
	return data
}
