package controllers

import (
	"example/project/Models"
	"example/project/initializers"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create post
	post := Models.Post{Title: body.Title, Body: body.Body}

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

func PostsIndex(c *gin.Context) {
	//get posts
	var posts []Models.Post
	initializers.DB.Find(&posts)

	//respons with that
	c.JSON(200, gin.H{"posts": posts})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	var post Models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{"post": post})
}

func PostUpdate(c *gin.Context) {
	//get the id of the url
	id := c.Param("id")

	//get the datat form req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post
	var post Models.Post
	initializers.DB.First(&post, id)

	//update it
	// post.Title = body.Title
	// post.Body = body.Body
	// initializers.DB.Save(&post)

	initializers.DB.Model(&post).Updates(Models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//responds with it
	c.JSON(200, gin.H{"post": post})
}

func PostDelet(c *gin.Context) {
	//get the post
	id := c.Param("id")

	//delete the post
	initializers.DB.Delete(&Models.Post{}, id)

	//give the response
	c.JSON(200, gin.H{"message": "Post deleted"})
}
