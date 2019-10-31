package routes

import "github.com/gin-gonic/gin"

func AddUsersRoutes(baseRouter gin.IRouter){
	v1 := baseRouter.Group("/api/v1/user")
	{
		v1.GET("/:id",fetchUser)
	}
	return
}

func fetchUser(ctx *gin.Context){

}