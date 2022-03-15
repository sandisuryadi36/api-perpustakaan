package config

import (
	"log"
	"os"
	"perpustakaan/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	// get database config
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	log.Println("env: ", user, password, host, port, database)
	
	// create postgres database connection
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
			&models.User{},
			&models.UserType{},
			&models.Book{},
			&models.Borrow{},
	)

	// insert default user type list
		var userType models.UserType
		init := db.Find(&userType)
		if init.RowsAffected == 0 {
			prim := []models.UserType{
				{Name: "admin", Description: "authorize as admin"},
				{Name: "basicUser", Description: "authorize as user"},
			}
			db.Create(&prim)
		}

	return db
}