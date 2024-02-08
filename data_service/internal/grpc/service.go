package grpcservice

import (
	"context"
	"fmt"

	pb "github.com/pavlechko/library/bookproto"
	"github.com/pavlechko/library/data_service/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IBook interface {
	GetBook(ctx context.Context, author, title string) (models.Book, error)
}

type serverAPI struct {
	pb.UnimplementedBookServiceServer
	book IBook
}

func Register(gRPC *grpc.Server, book IBook) {
	pb.RegisterBookServiceServer(gRPC, &serverAPI{book: book})
}

func (s *serverAPI) GetBookByAuthorAndTitle(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	if req.GetAuthor() == "" || req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Author and title are required")
	}

	book, err := s.book.GetBook(ctx, req.GetAuthor(), req.GetTitle())
	if err != nil {
		return nil, fmt.Errorf("Filed to get book")
	}

	return &pb.BookResponse{
		Book: &pb.Book{
			Title:   book.Title,
			Author:  book.Author,
			Country: book.Country,
		},
	}, nil
}
