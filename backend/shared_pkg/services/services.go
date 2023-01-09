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
	userClient := user.NewClient(os.Getenv("USER_GRPC_ADDR"))
	movieClient := movie.NewClient(os.Getenv("MOVIE_GRPC_ADDR"))
	dictionaryClient := dictionary.NewClient(os.Getenv("DICTIONARY_GRPC_ADDR"))
	verificationClient := verification.NewClient(os.Getenv("VERIFICATION_GRPC_ADDR"))
	vocabularyClient := vocabulary.NewClient(os.Getenv("VOCABULARY_GRPC_ADDR"), os.Getenv("KAFKA_ADDR"))

	return &Services{
		userClient,
		movieClient,
		dictionaryClient,
		verificationClient,
		vocabularyClient,
	}
}
