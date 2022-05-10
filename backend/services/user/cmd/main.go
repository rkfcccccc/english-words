package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/rkfcccccc/english_words/proto/users"
	"github.com/rkfcccccc/english_words/user/pkg/postgres"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("USER_SERVICE_DB"),
	)

	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	log.Printf("listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
