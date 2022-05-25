package dictionary_test

// NOTE: MongoDB must be running to run these tests

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/shared_pkg/mongodb"
	"github.com/stretchr/testify/assert"
)

func getRepository() (dictionary.Repository, error) {
	if err := godotenv.Load("../../../../.env"); err != nil {
		return nil, fmt.Errorf("godotenv.Load: %v", err)
	}

	db, err := mongodb.NewClient(
		context.Background(), os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"),
		"localhost", os.Getenv("MONGO_PORT"), os.Getenv("MONGO_DB"),
	)

	if err != nil {
		return nil, fmt.Errorf("mongodb.NewClient: %v", err)
	}

	collection := db.Collection("dictionary")
	return dictionary.NewMongoRepository(collection), nil
}

func generateRandomString() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

var repo, repoErr = getRepository()

func TestCreateWordIndex(t *testing.T) {
	if repoErr != nil {
		t.Fatalf("No repository: %v", repoErr)
	}

	err := repo.CreateWordIndex(context.Background())
	assert.Nil(t, err)
}

func TestGetById(t *testing.T) {
	if repoErr != nil {
		t.Fatalf("No repository: %v", repoErr)
	}

	entry, err := repo.GetById(context.Background(), "an invalid id string")
	assert.NotNil(t, err)
	assert.Nil(t, entry)

	entry, err = repo.GetById(context.Background(), "000000000000000000000000")
	assert.Nil(t, err)
	assert.Nil(t, entry)
}

func TestGetByWord(t *testing.T) {
	if repoErr != nil {
		t.Fatalf("No repository: %v", repoErr)
	}

	entry, err := repo.GetByWord(context.Background(), "word that is not exists")
	assert.Nil(t, entry)
	assert.Nil(t, err)
}

func TestScenario(t *testing.T) {
	if repoErr != nil {
		t.Fatalf("No repository: %v", repoErr)
	}

	word := fmt.Sprintf("test_%s", generateRandomString())
	expectedEntry := dictionary.WordEntry{
		Word:     word,
		Phonetic: "testing purpose word's phonetic",
		Meanings: []dictionary.Meaning{
			{
				PartOfSpeech: "noun",
				Definitions: []dictionary.Definition{
					{
						Text:    "some definition of word",
						Example: "example of that definition",
					},
				},
				Synonyms: []string{"synonym1", "synonym2"},
				Antonyms: []string{},
			},
			{
				PartOfSpeech: "verb",
				Definitions: []dictionary.Definition{
					{
						Text:    "some other definition of word",
						Example: "one more example of that definition",
					},
				},
				Synonyms: []string{},
				Antonyms: []string{},
			},
		},
	}

	wordId, err := repo.Create(context.Background(), &expectedEntry)
	assert.Nil(t, err)

	expectedEntry.Id = wordId

	actualEntry, err := repo.GetById(context.Background(), wordId)
	assert.Nil(t, err)
	assert.Equal(t, &expectedEntry, actualEntry, "retrieved entry should be equal to initial")

	actualEntry, err = repo.GetByWord(context.Background(), word)
	assert.Nil(t, err)
	assert.Equal(t, &expectedEntry, actualEntry, "retrieved entry should be equal to initial")

	err = repo.Delete(context.Background(), wordId)
	assert.Nil(t, err)

	entry, err := repo.GetById(context.Background(), wordId)
	assert.Nil(t, err)
	assert.Nil(t, entry, "retrieved entry should be nil because it was deleted")
}
