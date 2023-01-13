package reversoapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type client struct{}

func (*client) GetTranslation(ctx context.Context, input string, from string, to string) (*translationResponse, error) {
	url := "https://api.reverso.net/translate/v1/translation"

	requestBody, err := json.Marshal(translationRequestBody{
		Format: "text",
		From:   from, To: to,
		Input: input,

		Options: translationOptions{
			ContextResults:    true,
			LanguageDetection: false,
			SentenceSplitter:  false,

			Origin: "translation.web",
		},
	})

	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15")

	response, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("http.Post: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}

	fmt.Println(string(body))
	var result translationResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return &result, nil
}

func NewClient() Client {
	return &client{}
}
