package dictionary

// TODO: maybe again somehow get rid of duplicated structs
type Definition struct {
	Text    string `json:"text"`
	Example string `json:"example"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type WordEntry struct {
	Id       string    `bson:"_id,omitempty"`
	Word     string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Meanings []Meaning `json:"meanings"`
}
