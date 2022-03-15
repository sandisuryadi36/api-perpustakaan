package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"perpustakaan/models"
	"perpustakaan/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		ID, err := token.ExtractTokenID(c)
		user := models.User{}
		db.Find(&user, ID)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		} else if user.UserTypeID != 1 {
			err = errors.New("this page just for authorized role")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)
		ID, err := token.ExtractTokenID(c)
		user := models.User{}
		db.Find(&user, ID)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		} else if user.UserTypeID == 1 {
			c.Next()
		} else if paramID,_ := strconv.Atoi(c.Param("id")); int(ID) != paramID {
			err = errors.New("this page just for authorized user")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
