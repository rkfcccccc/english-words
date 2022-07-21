package vocabulary

type Challenge struct {
	WordId       string
	LearningStep int
}

// TODO: maybe put this duplicated declarations to one place
type WordAction struct {
	UserId int    `json:"user_id"`
	WordId string `json:"word_id"`
	Add    bool   `json:"add"`
}
