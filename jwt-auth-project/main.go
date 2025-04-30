package main

import (
	"jwt-auth/controllers"
	"jwt-auth/initializers"

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

	r.Run()
}
