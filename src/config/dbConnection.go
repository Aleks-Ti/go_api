package db

import (
	"api_fiber/src/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Crreating structure")
	godotenv.Load()
	db_host := os.Getenv("PG_HOST")
	db_user := os.Getenv("PG_USER")
	db_port := os.Getenv("PG_PORT")
	db_password := os.Getenv("PG_PASSWORD")
	db_name := os.Getenv("PG_DB_NAME")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", db_host, db_user, db_password, db_name, db_port)
	var dbConnection, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connection error")
	}

	DB = dbConnection
	fmt.Println("db connected successfully")

	AutoMigrate(dbConnection)
}

func AutoMigrate(dsn *gorm.DB) {
	dsn.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}
