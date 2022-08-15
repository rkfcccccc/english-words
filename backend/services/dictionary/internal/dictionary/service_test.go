package dictionary_test

import (
	"context"
	"testing"

	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/lemmatizer"
	"github.com/rkfcccccc/english_words/shared_pkg/dsync"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreate(t *testing.T) {
	repo := dictionary.NewRepositoryMock()
	stub := dsync.NewStubClient()
	dict := dictionaryapi.NewMockClient()
	lemm := lemmatizer.New("en")
	service := dictionary.NewService(repo, stub, dict, lemm)

	dEntry := dictionaryapi.Entry{
		Word:     "test entry word",
		Phonetic: "oajoadh",
		Meanings: []dictionaryapi.Meaning{
			{
				PartOfSpeech: "verb",
				Definitions: []dictionaryapi.Definition{
					{
						Text:    "some definition",
						Example: "some example",
					},
				},
				Synonyms: []string{},
				Antonyms: []string{"antonym of that word"},
			},
			{
				PartOfSpeech: "empty",
				Definitions:  []dictionaryapi.Definition{},
				Synonyms:     []string{},
				Antonyms:     []string{},
			},
		},
	}

	repo.PushEntryResponse(nil)
	repo.PushErrResponse(nil)

	dict.AddResponse(&dEntry, nil)

	repo.PushStringResponse("new_id")
	repo.PushErrResponse(nil)

	wordId, err := service.Create(context.Background(), "some word")

	assert.Nil(t, err)
	assert.Equal(t, wordId, "new_id")
}

func TestServiceGet(t *testing.T) {
	repo := dictionary.NewRepositoryMock()
	stub := dsync.NewStubClient()
	dict := dictionaryapi.NewMockClient()
	lemm := lemmatizer.New("en")
	service := dictionary.NewService(repo, stub, dict, lemm)

	expectedEntry := &dictionary.WordEntry{}

	// GetById
	repo.PushEntryResponse(expectedEntry)
	repo.PushErrResponse(nil)

	actualEntry, err := service.GetById(context.Background(), "some_id")

	assert.Nil(t, err)
	assert.Equal(t, expectedEntry, actualEntry)

	// GetByWord
	repo.PushEntryResponse(expectedEntry)
	repo.PushErrResponse(nil)

	actualEntry, err = service.GetByWord(context.Background(), "some_word")

	assert.Nil(t, err)
	assert.Equal(t, expectedEntry, actualEntry)
}

func TestDelete(t *testing.T) {
	repo := dictionary.NewRepositoryMock()
	stub := dsync.NewStubClient()
	dict := dictionaryapi.NewMockClient()
	lemm := lemmatizer.New("en")
	service := dictionary.NewService(repo, stub, dict, lemm)

	repo.PushErrResponse(nil)
	err := service.Delete(context.Background(), "some_id")

	assert.Nil(t, err)
}
