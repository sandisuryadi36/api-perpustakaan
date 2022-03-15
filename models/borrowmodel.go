package models

import "time"

type (
	//Borrow
	Borrow struct {
		ID         uint      `gorm:"primary_key" json:"id"`
		UserID     uint      `json:"userID"`
		BookID     uint      `json:"bookID"`
		BorrowDate string    `json:"borrowDate"`
		ReturnDate string    `json:"returnDate"`
		Status     string    `json:"status"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		User       User      `json:"-"`
		Book       Book      `json:"-"`
	}
)