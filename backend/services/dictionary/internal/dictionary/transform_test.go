package dictionary_test

import (
	"testing"

	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/stretchr/testify/assert"
)

func TestFromDictionaryApi(t *testing.T) {
	input := &dictionaryapi.Entry{
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
		},
	}

	expectedOutput := &dictionary.WordEntry{
		Word:     "test entry word",
		Phonetic: "oajoadh",
		Meanings: []dictionary.Meaning{
			{
				PartOfSpeech: "verb",
				Definitions: []dictionary.Definition{
					{
						Text:    "some definition",
						Example: "some example",
					},
				},
				Synonyms: []string{},
				Antonyms: []string{"antonym of that word"},
			},
		},
	}

	actualOutput := dictionary.TransformFromApi(input)
	assert.Equal(t, expectedOutput, actualOutput)
}
