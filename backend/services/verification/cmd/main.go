package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/rkfcccccc/english_words/services/verification/internal/verification"
	"github.com/rkfcccccc/english_words/services/verification/pkg/mail"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync/redsync"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
	)

	mailClient := mail.NewClient(
		os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"), os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"), "English App",
	)

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sync := redsync.NewClient(redis)

	repo := verification.NewPostgresRepository(db)
	service := verification.NewService(repo, sync, mailClient)
	server := verification.NewServer(service)

	listener, err := net.Listen("tcp", os.Getenv("VERIFICATION_GRPC_ADDR"))
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
