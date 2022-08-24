package dictionary_test

import (
	"testing"

	pb "github.com/rkfcccccc/english_words/proto/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/internal/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	"github.com/stretchr/testify/assert"

	. "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

func TestTransformFromApi(t *testing.T) {
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

	expectedOutput := &WordEntry{
		Word:     "test entry word",
		Phonetic: "oajoadh",
		Meanings: []Meaning{
			{
				PartOfSpeech: "verb",
				Definitions: []Definition{
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

func TestTransformToGRPC(t *testing.T) {
	input := &WordEntry{
		Word:     "test entry word",
		Phonetic: "oajoadh",
		Meanings: []Meaning{
			{
				PartOfSpeech: "verb",
				Definitions: []Definition{
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

	expectedOutput := &pb.WordEntry{
		Word:     "test entry word",
		Phonetic: "oajoadh",
		Meanings: []*pb.Meaning{
			{
				PartOfSpeech: "verb",
				Definitions: []*pb.Definition{
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

	actualOutput := dictionary.TransformToGRPC(input)
	assert.Equal(t, expectedOutput, actualOutput)
}
