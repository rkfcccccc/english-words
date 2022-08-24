package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/lemmatizer"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync/redsync"
	"github.com/rkfcccccc/english_words/shared_pkg/mongodb"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	db, err := mongodb.NewClient(
		context.Background(), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"),
	)

	if err != nil {
		log.Fatalf("failed to connect mongodb: %v", err)
	}

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	sync := redsync.NewClient(redis)
	dict := dictionaryapi.NewClient()
	lemm := lemmatizer.New("en")

	producer := dictionary.NewKafkaProducer()
	repo := dictionary.NewMongoRepository(db.Collection("dictionary"))
	service := dictionary.NewService(repo, sync, dict, lemm, producer)
	server := dictionary.NewServer(service)

	listener, err := net.Listen("tcp", os.Getenv("DICTIONARY_GRPC_ADDR"))
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
