package controllers

import (
	"fmt"
	"net/http"
	"perpustakaan/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	loginInput struct {
		Name    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	registerInput struct {
		Name       string `json:"name" binding:"required"`
		Password   string `json:"password" binding:"required"`
		UserTypeID uint   `json:"user_type_id" binding:"required"`
	}

	loginResponse struct {
		Message string `json:"message"`
		UserName	string `json:"user"`
		Token   string `json:"token"`
	}

	registerResponse struct {
		Message string `json:"message"`
		UserName	string `json:"user"`
	}
)

// LoginUser godoc
// @Summary Login as user.
// @Description Logging in to get jwt token to access api by roles.
// @Tags Login/Register
// @Param Body body loginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} loginResponse
// @Router /login [post]
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Name = input.Name
	u.Password = input.Password

	token, err := models.LoginCheck(u.Name, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "name or password is incorrect."})
		return
	}

	response := loginResponse{
		Message: "login success",
		UserName: u.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, response)

}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Login/Register
// @Param Body body registerInput true "the body to register a user"
// @Produce json
// @Success 200 {object} registerResponse
// @Router /register [post]
func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input registerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if input.UserTypeID != 1 && input.UserTypeID != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "use user_type_id 1 or 2"})
		return
	}

	u := models.User{}

	u.Name = input.Name
	u.Password = input.Password
	u.UserTypeID = input.UserTypeID

	_, err := u.SaveUser(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	response := registerResponse{
		Message: "Register success",
		UserName: u.Name,
	}

	c.JSON(http.StatusOK, response)

}