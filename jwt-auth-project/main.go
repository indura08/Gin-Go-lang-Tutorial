package main

import (
	"jwt-auth/controllers"
	"jwt-auth/initializers"
	"jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadDBConnection()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Hello from JWT auth project using gin and go",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

//test karaddi ghnna one curl commands:

//1. curl -X POST http://localhost:8080/login \
//   -H "Content-Type: application/json" \
//   -c cookies.txt \
//   -d '{"email":"you@example.com", "password":"yourpassword"}'
