package vocabulary

type WordData struct {
	UserId int    `db:"user_id"`
	WordId string `db:"word_id"`

	AlreadyLearned bool `db:"already_learned"`
	LearningStep   int  `db:"learning_step"`

	NextChallenge *int `db:"next_challenge"`

	Count int `db:"count"`
}

func (w *WordData) Changed() bool {
	return w.AlreadyLearned || w.LearningStep != 0 || w.NextChallenge != nil
}

type WordMessage struct {
	UserId int    `json:"user_id"`
	WordId string `json:"word_id"`
	Add    bool   `json:"add"`
}
