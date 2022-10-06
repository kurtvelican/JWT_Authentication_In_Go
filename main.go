package main

import (
	"JWT_Authentication_In_Go/controllers"
	"JWT_Authentication_In_Go/initializers"
	"JWT_Authentication_In_Go/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	err := r.Run()
	if err != nil {
		return
	}
}
