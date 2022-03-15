package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"perpustakaan/controllers"
	"perpustakaan/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/book/", controllers.GetAllBook)
	r.GET("/book/:id", controllers.GetBookByID)
	r.POST("/book", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware() , controllers.AddBook)
	r.PUT("/book/:id", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware() , controllers.UpdateBook)
	r.DELETE("/book/:id", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware() , controllers.DeleteBook)

	r.POST("/borrow", middlewares.JwtAuthMiddleware(), controllers.AddBorrow)
	r.PUT("/return", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware(), controllers.ReturnBook)
	r.GET("/borrow", middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware(), controllers.GetBorrowList)
	r.GET("/borrow/user/:id", middlewares.JwtAuthMiddleware(), middlewares.UserMiddleware(), controllers.GetBorrowListByUserID)

	return r
}