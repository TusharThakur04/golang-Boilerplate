package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tusharthakur04/golang-CRUD/routes"
)

func main() {
	r := gin.Default()

	userGroup  := r.Group("/users")

	routes.RegisterUserRoutes(userGroup)


	r.Run(":8000")
}
