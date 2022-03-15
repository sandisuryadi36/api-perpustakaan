package models

type (
	// UserType
	UserType struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		User        []User    `json:"-"`
	}
)
