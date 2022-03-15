package controllers

import (
	"net/http"
	"perpustakaan/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	borrowInput struct {
		UserID     uint      `json:"userID"`
		BookID     uint      `json:"bookID"`
		Days	   int      `json:"days"`
	}

	returnInput struct {
		UserID     uint      `json:"userID"`
		BookID     uint      `json:"bookID"`
	}

	borrowResponse struct {
		Message string `json:"message"`
		BorrowStatus string `json:"borrowStatus"`
		ReturnDate string `json:"returnDate"`
		Book models.Book `json:"book"`
	}
)

// Add borrow list
// AddBorrowList godoc
// @Summary Add new borrow.
// @Description Add new borrow to borrow list.
// @Tags Borrow
// @Param Body body borrowInput true "the body to add a borrow"
// @Produce json
// @Success 200 {object} borrowResponse
// @Router /borrow [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func AddBorrow(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input borrowInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrow := models.Borrow{}
	check := db.Where("user_id = ? AND book_id = ? AND status = 'Borrowed'", input.UserID, input.BookID).First(&borrow)
	if check.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book already borrowed"})
		return
	}
	
	borrow.UserID = input.UserID
	borrow.BookID = input.BookID
	borrow.BorrowDate = time.Now().Format("2006-01-02")
	borrow.ReturnDate = time.Now().AddDate(0, 0, input.Days).Format("2006-01-02")
	borrow.Status = "Borrowed"

	tx := db.Begin()
	if err:= tx.Create(&borrow).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	book := models.Book{}
	tx.Find(&book, input.BookID)
	book.Stock = book.Stock - 1
	if err:= tx.Model(&book).Update("stock", book.Stock).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	// join table by book id
	db.Joins("JOIN borrows ON borrows.book_id = books.id").Where("borrows.book_id = ?", input.BookID).Find(&book)

	response := borrowResponse{
		Message: "Success add borrow",
		BorrowStatus: borrow.Status,
		ReturnDate: borrow.ReturnDate,
		Book: book,
	}
	c.JSON(http.StatusOK, response)
}

// Return book
// ReturnBook godoc
// @Summary Return book.
// @Description Return book.
// @Tags Borrow
// @Param Body body returnInput true "the body to return a book"
// @Produce json
// @Success 200 {object} borrowResponse
// @Router /return [put]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func ReturnBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input returnInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrow := models.Borrow{}
	if err := db.Where("user_id = ? AND book_id = ?", input.UserID, input.BookID).First(&borrow).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrow.ReturnDate = time.Now().Format("2006-01-02")
	borrow.Status = "Returned"

	tx := db.Begin()
	if err:= tx.Save(&borrow).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{}
	tx.Find(&book, input.BookID)
	book.Stock = book.Stock + 1
	if err:= tx.Model(&book).Update("stock", book.Stock).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx.Commit()

	// join table by book id
	db.Joins("JOIN borrows ON borrows.book_id = books.id").Where("borrows.book_id = ?", input.BookID).Find(&book)

	response := borrowResponse{
		Message: "Success return book",
		BorrowStatus: borrow.Status,
		ReturnDate: borrow.ReturnDate,
		Book: book,
	}
	c.JSON(http.StatusOK, response)
}

// GET borrow list
// GetBorrowList godoc
// @Summary Get borrow list.
// @Description Get borrow list.
// @Tags Borrow
// @Produce json
// @Success 200 {object} []models.Borrow
// @Router /borrow [get]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func GetBorrowList(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var borrows []models.Borrow
	db.Find(&borrows)

	c.JSON(http.StatusOK, borrows)
}

// GET borrow list by user id
// GetBorrowListByUserID godoc
// @Summary Get borrow list by user id.
// @Description Get borrow list by user id.
// @Tags Borrow
// @Produce json
// @Success 200 {object} []models.Borrow
// @Router /borrow/user/{user_id} [get]
// @Param user_id path string true "user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func GetBorrowListByUserID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var borrows []models.Borrow
	db.Where("user_id = ?", c.Param("id")).Find(&borrows)

	c.JSON(http.StatusOK, borrows)
}