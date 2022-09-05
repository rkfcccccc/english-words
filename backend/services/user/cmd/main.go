package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/rkfcccccc/english_words/services/user/internal/user"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync/redsync"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"google.golang.org/grpc"
)

func main() {
	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
	)

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sync := redsync.NewClient(redis)

	repo := user.NewPostgresRepository(db)
	service := user.NewService(repo, sync)
	server := user.NewServer(service)

	listener, err := net.Listen("tcp", os.Getenv("USER_GRPC_ADDR"))
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
