package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/rkfcccccc/english_words/services/movie/internal/movie"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync/redsync"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"github.com/rkfcccccc/english_words/shared_pkg/services/dictionary"
	"github.com/rkfcccccc/english_words/shared_pkg/services/vocabulary"
	"google.golang.org/grpc"
)

func main() {
	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
	)

	dictionary := dictionary.NewClient("localhost" + os.Getenv("DICTIONARY_GRPC_ADDR"))
	vocabulary := vocabulary.NewClient("localhost"+os.Getenv("VOCABULARY_GRPC_ADDR"), os.Getenv("KAFKA_ADDR"))

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sync := redsync.NewClient(redis)

	repo := movie.NewPostgresRepository(db)
	service := movie.NewService(repo, sync, dictionary, vocabulary)
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
