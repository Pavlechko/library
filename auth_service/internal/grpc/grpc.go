package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/pavlechko/library/bookproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// type IBook interface {
// 	GetBookByAuthorAndTitle(ctx context.Context, author, title string) (models.Book, error)
// }

type GrpcClient struct {
	client pb.BookServiceClient
	// book   IBook
}

func NewClient() GrpcClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewBookServiceClient(conn)
	return GrpcClient{client: client}
}

func (c *GrpcClient) GetBookByAuthorAndTitle(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	if req.GetAuthor() == "" || req.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Author and title are required")
	}
	resp, err := c.client.GetBookByAuthorAndTitle(ctx, &pb.BookRequest{Author: req.GetAuthor(), Title: req.GetTitle()})
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}
