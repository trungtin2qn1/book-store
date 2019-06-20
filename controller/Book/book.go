package Book

import (
	"book-store/database"
	"book-store/model/Book"
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetBooks ...
func GetBooks(c *gin.Context) {
	option := c.Query("option")
	books := []database.Book{}
	if option == "search" {
		query := c.Query("query")
		books, _ = Book.SearchBooks(query)
	} else {
		books, _ = Book.GetAllBooks()
	}
	c.JSON(200, books)
}

//GetBooksInfoByID ...
func GetBookInfoByID(c *gin.Context) {
	bookID := c.Param("book_id")
	fmt.Println("Book ID:", bookID)
	book, _ := Book.GetBookByID(bookID)
	c.JSON(200, book)
}

//CreateBook ...
func CreateBook(c *gin.Context) {
	book := database.Book{}
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Data or data type is invalid",
		})
		return
	}
	fmt.Println("Book:", book)
	fmt.Println("Book[images]", book.Images)
	res, err := Book.CreateBook(&book)
	if err != nil {
		c.JSON(406, gin.H{
			"message": "Check input again",
		})
		return
	}

	c.JSON(200, res)
}

//UpdateBookInfo ...
func UpdateBookInfo(c *gin.Context) {

}

//DeleteBook ...
func DeleteBook(c *gin.Context) {

}
