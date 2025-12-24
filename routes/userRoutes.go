package routes

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.RouterGroup){

	router.GET("/" ,func(c *gin.Context){
		c.JSON(200 , gin.H{
			"users":"user list",
		})

	})
	router.GET("/:id", Getuser)
	router.POST("/", CreateUser)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DeleteUser)

}


