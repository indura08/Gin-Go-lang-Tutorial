package controllers

import (
	"example/project/Models"
	"example/project/initializers"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data off request body

	//create post
	post := Models.Post{Title: "Hello", Body: "Post body"}

	result := initializers.DB.Create(&post)

	//return it
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
