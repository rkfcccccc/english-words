package dictionary

import (
	pb "github.com/rkfcccccc/english_words/proto/dictionary"

	. "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

func transformFromGRPC(dEntry *pb.WordEntry) *WordEntry {
	if dEntry == nil {
		return nil
	}

	entry := WordEntry{
		Id:           dEntry.Id,
		Word:         dEntry.Word,
		Phonetic:     dEntry.Phonetic,
		Translations: dEntry.Translations,
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

	if dEntry.Pictures != nil {
		entry.Pictures = make([]SourcedPicture, len(dEntry.Pictures))
		for i, picture := range dEntry.Pictures {
			entry.Pictures[i] = SourcedPicture{
				Url:    picture.Url,
				Source: picture.Source,
			}
		}
	}

	return &entry
}
