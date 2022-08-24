package dictionary

type SourcedPicture struct {
	Url    string `json:"url"`
	Source string `json:"source"`
}

type Definition struct {
	Text    string `json:"text"`
	Example string `json:"example"`
}

type Meaning struct {
	PartOfSpeech string       `json:"part_of_speech"`
	Definitions  []Definition `json:"definitions"`
	Synonyms     []string     `json:"synonyms"`
	Antonyms     []string     `json:"antonyms"`
}

type WordEntry struct {
	Id       string    `bson:"_id,omitempty"`
	Word     string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Meanings []Meaning `json:"meanings"`

	Pictures []SourcedPicture `json:"pictures"`
}
