package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/vocabulary/internal/vocabulary"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
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

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	state := vocabulary.NewRedisVocabularyState(redis)
	repo := vocabulary.NewPostgresRepository(db)

	service := vocabulary.NewService(repo, state)
	server := vocabulary.NewServer(service)

	listener, err := net.Listen("tcp", os.Getenv("VOCABULARY_GRPC_ADDR"))
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
