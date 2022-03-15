package models

import (
	"time"

	"perpustakaan/utils/token"

	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type (
	// User
	User struct {
		ID         uint      `gorm:"primary_key" json:"id"`
		Name       string    `json:"name"`
		Password   string    `json:"password"`
		UserTypeID uint      `json:"userTypeID"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		UserType   UserType  `json:"-"`
		Borrow     []Borrow  `json:"-"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(name string, password string, db *gorm.DB) (string, error) {

	var err error
	u := User{}

	err = db.Model(User{}).Where("name = ?", name).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}