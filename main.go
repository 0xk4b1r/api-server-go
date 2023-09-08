package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Bhagavad Gita", Author: "Vedavyasa", Quantity: 10},
	{ID: "2", Title: "Ramayana", Author: "Valmiki", Quantity: 10},
	{ID: "3", Title: "Rigveda", Author: "Various", Quantity: 5},
	{ID: "4", Title: "Yajurveda", Author: "Various", Quantity: 6},
	{ID: "5", Title: "Samaveda", Author: "Various", Quantity: 7},
	{ID: "6", Title: "Atharvaveda", Author: "Various", Quantity: 8},
	{ID: "7", Title: "Upanishads", Author: "Various", Quantity: 9},
}

// ------------------------------------ //

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")

}

func addBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	
}

func main() {
	router := gin.Default()

	router.GET("/")
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", addBook)

	router.Run("localhost:8080")
}
