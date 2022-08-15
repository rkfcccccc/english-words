package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/movie/internal/movie"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/services/dictionary"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
	)

	dict := dictionary.NewClient("localhost" + os.Getenv("DICTIONARY_GRPC_ADDR"))

	repo := movie.NewPostgresRepository(db)
	service := movie.NewService(repo, dict)
	server := movie.NewServer(service)

	listener, err := net.Listen("tcp", os.Getenv("MOVIE_GRPC_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	server.Register(s)

	log.Printf("listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
