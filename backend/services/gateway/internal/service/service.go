package service

import (
	"os"

	"github.com/rkfcccccc/english_words/services/gateway/internal/service/dictionary"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/movie"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/user"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/verification"
	"github.com/rkfcccccc/english_words/services/gateway/internal/service/vocabulary"
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
	vocabularyClient := vocabulary.NewClient("localhost"+os.Getenv("VOCABULARY_GRPC_ADDR"), "localhost:9092")

	return &Services{
		userClient,
		movieClient,
		dictionaryClient,
		verificationClient,
		vocabularyClient,
	}
}
