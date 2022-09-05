package dictionary_test

// NOTE: MongoDB must be running to run these tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/shared_pkg/mongodb"
	"github.com/stretchr/testify/assert"

	. "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

func getRepository() (dictionary.Repository, error) {
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

	word := "test_word"
	expectedEntry := WordEntry{
		Word:     word,
		Phonetic: "testing purpose word's phonetic",
		Meanings: []Meaning{
			{
				PartOfSpeech: "noun",
				Definitions: []Definition{
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
				Definitions: []Definition{
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

	expectedEntry.Pictures = []SourcedPicture{{Url: "picture1.png", Source: "source1"}, {Url: "picture2.png", Source: "source2"}}

	err = repo.SetPictures(context.Background(), wordId, expectedEntry.Pictures)
	assert.Nil(t, err)

	actualEntry, err = repo.GetByWord(context.Background(), word)
	assert.Nil(t, err)
	assert.Equal(t, &expectedEntry, actualEntry, "retrieved entry should be equal to initial")

	err = repo.Delete(context.Background(), wordId)
	assert.Nil(t, err)

	entry, err := repo.GetById(context.Background(), wordId)
	assert.Nil(t, err)
	assert.Nil(t, entry, "retrieved entry should be nil because it was deleted")
}

func TestSearch(t *testing.T) {
	wordIds := []string{}

	for i := 0; i < 3; i++ {
		entry := &WordEntry{Word: fmt.Sprintf("test_search_%d", i)}
		wordId, err := repo.Create(context.Background(), entry)
		assert.Nil(t, err)

		wordIds = append(wordIds, wordId)
	}

	results, err := repo.Search(context.Background(), "test_search_")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(results))

	for i := 0; i < 3; i++ {
		err := repo.Delete(context.Background(), wordIds[i])
		assert.Nil(t, err)
	}
}
