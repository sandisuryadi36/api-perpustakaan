package main

import (
	"log"
	"os"
	"perpustakaan/config"
	"perpustakaan/docs"
	"perpustakaan/routes"
	"perpustakaan/utils"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

// @contact.name Sandi Suryadi
// @contact.email sandisuryadi.mail@gmail.com

func main() {
	// load env
	evirontment := utils.Getenv("ENVIRONMENT", "development")

	if evirontment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Perpustakaan API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// router
	r := routes.SetupRouter(db)
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.Run()
}