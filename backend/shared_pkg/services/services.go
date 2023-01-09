package services

import (
	"os"

	"github.com/rkfcccccc/english_words/shared_pkg/services/dictionary"
	"github.com/rkfcccccc/english_words/shared_pkg/services/movie"
	"github.com/rkfcccccc/english_words/shared_pkg/services/user"
	"github.com/rkfcccccc/english_words/shared_pkg/services/verification"
	"github.com/rkfcccccc/english_words/shared_pkg/services/vocabulary"
)

type Services struct {
	User         *user.Client
	Movie        *movie.Client
	Dictionary   *dictionary.Client
	Verification *verification.Client
	Vocabulary   *vocabulary.Client
}

func NewService() *Services {
	userClient := user.NewClient("localhost" + os.Getenv("USER_GRPC_ADDR"))
	movieClient := movie.NewClient("localhost" + os.Getenv("MOVIE_GRPC_ADDR"))
	dictionaryClient := dictionary.NewClient("localhost" + os.Getenv("DICTIONARY_GRPC_ADDR"))
	verificationClient := verification.NewClient("localhost" + os.Getenv("VERIFICATION_GRPC_ADDR"))
	vocabularyClient := vocabulary.NewClient("localhost"+os.Getenv("VOCABULARY_GRPC_ADDR"), os.Getenv("KAFKA_ADDR"))

	return &Services{
		userClient,
		movieClient,
		dictionaryClient,
		verificationClient,
		vocabularyClient,
	}
}
