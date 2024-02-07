package grpcservice

import (
	"context"

	pb "github.com/pavlechko/library/bookproto"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedBookServiceServer
}

func Register(gRPC *grpc.Server) {
	pb.RegisterBookServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetBookByAuthorAndTitle(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{
		Book: &pb.Book{
			Title:   req.GetTitle(),
			Author:  "Some author",
			Country: "USA",
		},
	}, nil
}
