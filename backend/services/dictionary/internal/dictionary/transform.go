package dictionary

import (
	pb "github.com/rkfcccccc/english_words/proto/dictionary"
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
	models "github.com/rkfcccccc/english_words/shared_pkg/services/dictionary/models"
)

func TransformFromApi(dEntry *dictionaryapi.Entry) *models.WordEntry {
	entry := models.WordEntry{
		Word:     dEntry.Word,
		Phonetic: dEntry.Phonetic,
	}

	entry.Meanings = make([]models.Meaning, len(dEntry.Meanings))
	for i, meaning := range dEntry.Meanings {
		entry.Meanings[i] = models.Meaning{
			PartOfSpeech: meaning.PartOfSpeech,
			Synonyms:     meaning.Synonyms,
			Antonyms:     meaning.Antonyms,
		}

		entry.Meanings[i].Definitions = make([]models.Definition, len(meaning.Definitions))
		for j, definition := range meaning.Definitions {
			entry.Meanings[i].Definitions[j] = models.Definition{
				Text:    definition.Text,
				Example: definition.Example,
			}
		}
	}

	return &entry
}

func TransformToGRPC(dEntry *models.WordEntry) *pb.WordEntry {
	if dEntry == nil {
		return nil
	}

	entry := pb.WordEntry{
		Id:           dEntry.Id,
		Word:         dEntry.Word,
		Phonetic:     dEntry.Phonetic,
		Translations: dEntry.Translations,
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

	if dEntry.Pictures != nil {
		entry.Pictures = make([]*pb.SourcedPicture, len(dEntry.Pictures))
		for i, picture := range dEntry.Pictures {
			entry.Pictures[i] = &pb.SourcedPicture{
				Url:    picture.Url,
				Source: picture.Source,
			}
		}
	}

	return &entry
}
