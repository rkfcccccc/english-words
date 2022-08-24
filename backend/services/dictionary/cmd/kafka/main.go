package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi/freepik"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/pictureapi/unsplash"
	"github.com/rkfcccccc/english_words/shared_pkg/mongodb"
	"github.com/segmentio/kafka-go"
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

	pictures := pictureapi.NewService(freepik.NewPictureRepository(), unsplash.NewPictureRepository())

	repo := dictionary.NewMongoRepository(db.Collection("dictionary"))
	service := dictionary.NewService(repo, nil, nil, nil, nil)

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_ADDRESS"), "pictures", 0)
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}

	consumer := dictionary.NewKafkaConsumer(service, pictures)
	if err := consumer.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
