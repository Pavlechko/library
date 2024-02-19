package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pavlechko/library/auth_service/internal/grpc"
	"github.com/pavlechko/library/auth_service/internal/models"
	pb "github.com/pavlechko/library/bookproto"
	amqp "github.com/rabbitmq/amqp091-go"
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
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
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

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"author",
			false,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "Failed to declare a queue")

		body := book.Author

		err = ch.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)

		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)

		msgs, err := ch.Consume(
			q.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		failOnError(err, "Failed to register a consumer")

		aut := make(chan string)

		go func() {
			for d := range msgs {
				aut <- string(d.Body)
			}

		}()
		res, err := c.service.GetBookByAuthorAndTitle(ctx, &pb.BookRequest{Author: <-aut, Title: book.Title})
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
