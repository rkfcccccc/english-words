package reversoapi

type translationOptions struct {
	SentenceSplitter  bool `json:"sentenceSplitter"`
	ContextResults    bool `json:"contextResults"`
	LanguageDetection bool `json:"languageDetection"`

	Origin string `json:"origin"`
}

type translationRequestBody struct {
	Format string `json:"format"`

	Input string `json:"input"`
	From  string `json:"from"`
	To    string `json:"to"`

	Options translationOptions `json:"options"`
}

type translationResponse struct {
	ContextResults *struct {
		Results []*struct {
			Translation string `json:"translation"`
			Frequency   int    `json:"frequency"`
		}
	} `json:"contextResults"`
}
