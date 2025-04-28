package main

import (
	"example/project/Models"
	"example/project/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&Models.Post{})
}
