package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/joho/godotenv/autoload"
)

func ConnectToDB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(os.Getenv("DB_URL")))
}
