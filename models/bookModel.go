package models

import "time"

type (
	// Book
	Book struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Author      string    `json:"author"`
		Publisher   string    `json:"publisher"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Borrow      []Borrow  `json:"-"`
	}
)