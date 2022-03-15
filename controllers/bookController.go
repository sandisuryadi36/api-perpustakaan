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

	getResponse struct {
		Message string `json:"message"`
		Books []models.Book `json:"books"`
	}

	getOneResponse struct {
		Message string `json:"message"`
		Books models.Book `json:"books"`
	}
)

// POST book
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

// PUT book
// UpdateBook godoc
// @Summary Edit book.
// @Description Edit a book in list.
// @Tags Book
// @Param id path string true "Book id"
// @Param Body body bookInput true "the body to edit a book"
// @Produce json
// @Success 200 {object} bookResponse
// @Router /book/{id} [PUT]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateBook := models.Book{}
	updateBook.Name = input.Name
	updateBook.Description = input.Description
	updateBook.Author = input.Author
	updateBook.Publisher = input.Publisher

	db.Model(&book).Updates(updateBook)

	response := bookResponse{
		Message: "Book has been updated",
		ID:          book.ID,
		Name:        book.Name,
	}
	c.JSON(http.StatusOK, response)
}

// GET all book
// GetAllBook godoc
// @Summary Get all book.
// @Description Get all book in list.
// @Tags Book
// @Produce json
// @Success 200 {object} getResponse
// @Router /book [get]
func GetAllBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)

	response := getResponse{
		Message: "Success",
		Books: books,
	}
	c.JSON(http.StatusOK, response)
}

// GET book by id
// GetBookByID godoc
// @Summary Get book by id.
// @Description Get a book in list.
// @Tags Book
// @Param id path string true "Book id"
// @Produce json
// @Success 200 {object} getOneResponse
// @Router /book/{id} [get]
func GetBookByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	response := getOneResponse{
		Message: "Success",
		Books: book,
	}
	c.JSON(http.StatusOK, response)
}

// DELETE book
// DeleteBook godoc
// @Summary Delete book.
// @Description Delete a book in list.
// @Tags Book
// @Param id path string true "Book id"
// @Produce json
// @Success 200 {object} bookResponse
// @Router /book/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	response := bookResponse{
		Message: "Book has been deleted",
		ID:          book.ID,
		Name:        book.Name,
	}
	c.JSON(http.StatusOK, response)
}