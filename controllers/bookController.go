package controllers

import (
	"net/http"
	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	bookInput struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Publisher   string `json:"publisher"`
	}

	bookResponse struct {
		Message string `json:"message"`
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
	}
)

// AddBook godoc
// @Summary Add new book.
// @Description Add new book to book list.
// @Tags Book
// @Param Body body bookInput true "the body to add a book"
// @Produce json
// @Success 200 {object} bookResponse
// @Router /book [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func AddBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{}
	book.Name = input.Name
	book.Description = input.Description
	book.Author = input.Author
	book.Publisher = input.Publisher

	db.Create(&book)

	response := bookResponse{
		Message: "Book has been created",
		ID:          book.ID,
		Name:        book.Name,
	}
	c.JSON(http.StatusOK, response)
}