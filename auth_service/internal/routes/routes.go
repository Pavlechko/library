package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavlechko/library/auth_service/internal/grpc"
	"github.com/pavlechko/library/auth_service/internal/models"
	pb "github.com/pavlechko/library/bookproto"
)

type APIClient struct {
	listenPort string
	service    grpc.GrpcClient
}

func NewAPIClient(port string, service grpc.GrpcClient) *APIClient {
	return &APIClient{
		listenPort: port,
		service:    service,
	}
}

func (c *APIClient) Run(ctx context.Context) {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var book models.BookInput
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		res, err := c.service.GetBookByAuthorAndTitle(ctx, &pb.BookRequest{Author: book.Author, Title: book.Title})
		if err != nil {
			http.Error(w, "Error Respons from gRPC", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res.Book.Country)
	})

	http.Handle("/", r)
	fmt.Println("Server is listening on port: ", c.listenPort)

	server := &http.Server{
		Addr: c.listenPort,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error listening server:", err.Error())
	}
}
