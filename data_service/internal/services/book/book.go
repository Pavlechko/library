package book

import (
	"context"
	"errors"
	"fmt"

	"github.com/pavlechko/library/data_service/internal/models"
	"github.com/pavlechko/library/data_service/internal/storage"
	"github.com/pavlechko/library/data_service/internal/utils"
)

type BookService struct {
	bookProvider BookProvider
}

type BookProvider interface {
	Book(ctx context.Context, author, title string) (models.Book, error)
}

func NewBookService(bookProvider BookProvider) *BookService {
	utils.InfoLogger.Println("NewBookService: ", bookProvider)
	return &BookService{
		bookProvider: bookProvider,
	}
}

func (b *BookService) GetBook(ctx context.Context, author, title string) (models.Book, error) {
	book, err := b.bookProvider.Book(ctx, author, title)
	if err != nil {
		if errors.Is(err, storage.ErrBookNotFound) {
			fmt.Println("Book not found")
			panic(err)

			// return book, fmt.Errorf("%w", errors.New("Invalid credentials"))
		}
		panic(err)

		// return book, fmt.Errorf("%w", errors.New("Unexpected error"))
	}

	return book, nil
}
