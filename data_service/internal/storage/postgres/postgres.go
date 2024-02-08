package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pavlechko/library/data_service/internal/models"
	"github.com/pavlechko/library/data_service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func ConnectDB() (*Storage, error) {
	utils.InfoLogger.Println("Connect DB")

	err := godotenv.Load()
	if err != nil {
		utils.ErrorLogger.Println("Error loading .env file: %w", err)
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("POSTGRES_PORT"),
	)

	utils.InfoLogger.Println("dsn: ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	utils.InfoLogger.Println("Error: ", err)

	if err != nil {
		utils.ErrorLogger.Println("Error reading HTTP response body:", err)
		return nil, err
	}
	utils.InfoLogger.Println("Connected successfully to the Database")

	return &Storage{db: db}, err
}

func (s *Storage) Close() {
	sqlDB, err := s.db.DB()
	if err != nil {
		fmt.Println("Error closing database connection:", err)
		return
	}
	sqlDB.Close()
}

func (s *Storage) Book(ctx context.Context, author, title string) (models.Book, error) {
	var bookModel models.Book
	res := s.db.Raw("SELECT id, author, title, country FROM books WHERE author=? AND title=?", author, title).Scan(&bookModel)

	if res.Error != nil {
		fmt.Println("Error finding user", res.Error.Error())
		// return bookModel, fmt.Errorf("%s", res.Error.Error())
		panic(res)
	}

	return bookModel, nil
}
