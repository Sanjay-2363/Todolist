package repository

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Connect to database
func DbConnection() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	Db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
