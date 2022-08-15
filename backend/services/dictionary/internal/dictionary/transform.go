package dictionary

import (
	pb "github.com/rkfcccccc/english_words/proto/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
)

func TransformFromApi(dEntry *dictionaryapi.Entry) *WordEntry {
	entry := WordEntry{
		Word:     dEntry.Word,
		Phonetic: dEntry.Phonetic,
	}

	entry.Meanings = make([]Meaning, len(dEntry.Meanings))
	for i, meaning := range dEntry.Meanings {
		entry.Meanings[i] = Meaning{
			PartOfSpeech: meaning.PartOfSpeech,
			Synonyms:     meaning.Synonyms,
			Antonyms:     meaning.Antonyms,
		}

		entry.Meanings[i].Definitions = make([]Definition, len(meaning.Definitions))
		for j, definition := range meaning.Definitions {
			entry.Meanings[i].Definitions[j] = Definition{
				Text:    definition.Text,
				Example: definition.Example,
			}
		}
	}

	return &entry
}

func TransformToGRPC(dEntry *WordEntry) *pb.WordEntry {
	if dEntry == nil {
		return nil
	}

	entry := pb.WordEntry{
		Word:     dEntry.Word,
		Phonetic: dEntry.Phonetic,
	}

	entry.Meanings = make([]*pb.Meaning, len(dEntry.Meanings))
	for i, meaning := range dEntry.Meanings {
		entry.Meanings[i] = &pb.Meaning{
			PartOfSpeech: meaning.PartOfSpeech,
			Synonyms:     meaning.Synonyms,
			Antonyms:     meaning.Antonyms,
		}

		entry.Meanings[i].Definitions = make([]*pb.Definition, len(meaning.Definitions))
		for j, definition := range meaning.Definitions {
			entry.Meanings[i].Definitions[j] = &pb.Definition{
				Text:    definition.Text,
				Example: definition.Example,
			}
		}
	}

	return &entry
}
