package main

import (
	"example/project/controllers"
	"example/project/initializers"

	"github.com/gin-gonic/gin"
)

func init() { // In Go, func init() is a special function that automatically runs before the main() function. You don't need to call init() manually — Go calls it for you.

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PostCreate)

	r.Run()
}
