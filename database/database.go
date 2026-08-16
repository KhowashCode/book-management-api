package database

import (
	"book-management-api/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	defer registerTable()

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Failed to load dotenv!")
	}

	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		port     = os.Getenv("DB_PORT")
		table    = os.Getenv("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, table, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connect To Database")
	}

	DB = db
}

func registerTable() {
	DB.AutoMigrate(&models.Book{}, &models.Category{})
}
