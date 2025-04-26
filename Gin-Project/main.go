package main

import (
	"errors"
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

func createBook(c *gin.Context) {
	var newBook Book
	if error := c.BindJSON(&newBook); error != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, "Book Created") //methan onnm newBook denna puluwan Book created wenuwat
}

func bookByIt(c *gin.Context) {
	id := c.Param("id") //pi/book/1 wage
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"}) //gin.H{}	Shortcut for map[string]interface{}
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func checkout(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not found"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Book is not available at the moment"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book checked out successfully", "book": book})
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No request param , check again"})
		return
	}

	book, error := getBookById(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book is not found , check again"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}

func main() {
	//1.setting up router which handles diffrent routes , endpoints of our application
	router := gin.Default()
	router.GET("/getBooks", GetAllBooks)
	router.POST("/create", createBook) //curl localhost:8080/create --include --header "Content-Type: application/json" -d @body.json --request POST menna me command ek ghla curl walin giyahki postman nathuwa
	router.GET("/getBook/:id", bookByIt)
	router.PATCH("/checkout", checkout)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
