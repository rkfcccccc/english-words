package dictionary

import (
	"github.com/rkfcccccc/english_words/services/dictionary/pkg/dictionaryapi"
)

func TransformFromApi(dEntry *dictionaryapi.Entry) *WordEntry {
	entry := WordEntry{
		Word:     dEntry.Word,
		Phonetic: dEntry.Phonetic,
	}

	entry.Meanings = make([]Meaning, len(dEntry.Meanings))
	for i, meaning := range dEntry.Meanings {
		entry.Meanings[i].PartOfSpeech = meaning.PartOfSpeech
		entry.Meanings[i].Synonyms = meaning.Synonyms
		entry.Meanings[i].Antonyms = meaning.Antonyms

		entry.Meanings[i].Definitions = make([]Definition, len(meaning.Definitions))
		for j, definition := range meaning.Definitions {
			entry.Meanings[i].Definitions[j].Text = definition.Text
			entry.Meanings[i].Definitions[j].Example = definition.Example
		}
	}

	return &entry
}
