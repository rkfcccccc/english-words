package main

import (
	"context"
	"log"
	"os"

	"github.com/rkfcccccc/english_words/services/vocabulary/internal/vocabulary"
	"github.com/rkfcccccc/english_words/shared_pkg/postgres"
	"github.com/rkfcccccc/english_words/shared_pkg/redis"
	"github.com/segmentio/kafka-go"
)

func main() {
	db := postgres.NewPool(
		context.Background(), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"),
	)

	redis := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	state := vocabulary.NewRedisVocabularyState(redis)
	repo := vocabulary.NewPostgresRepository(db)
	service := vocabulary.NewService(repo, state)

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_ADDR"), "vocabulary", 0)
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}

	consumer := vocabulary.NewKafkaConsumer(service, os.Getenv("KAFKA_ADDR"))
	if err := consumer.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
