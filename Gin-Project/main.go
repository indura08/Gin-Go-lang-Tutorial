package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json: "id"`
	Title    string `json: "titile"`
	Author   string `json: "author"`
	Quantity int    `json "quantity"`
	//methna me book struct eke feilds okkoma capital letters walin daanna one , ehma une nattnm me file ekn pita anik file walt serialize krgnna bari wena book struct ekt adla objectss
}

var books = []Book{
	{Id: "1", Title: "Fc Bayern Munchen: History", Author: "Franz Bechanbuer", Quantity: 10},
	{Id: "2", Title: "Fc Barcelona: History", Author: "Lionel Messi", Quantity: 5},
	{Id: "3", Title: "Real Madrid: History", Author: "Cristiano Ronaldo", Quantity: 7},
	{Id: "4", Title: "Manchester United: History", Author: "David Beckham", Quantity: 3},
	{Id: "5", Title: "Liverpool: History", Author: "Steven Gerrard", Quantity: 8},
	{Id: "6", Title: "Chelsea: History", Author: "Frank Lampard", Quantity: 4},
	{Id: "7", Title: "Arsenal: History", Author: "Thierry Henry", Quantity: 6},
}

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	//1.setting up router which handles diffrent routes , endpoints of our application
	router := gin.Default()
	router.GET("/getBooks", GetAllBooks)
	router.Run("localhost:8080")
}
