package dictionaryapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrNoDefinitionsFound = errors.New("no definitions found")

func fetchWordEntries(ctx context.Context, language, word string) ([]Entry, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", language, word)
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("http.Get: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, ErrNoDefinitionsFound
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("bad status %s with body %s", response.Status, bytes)
	}

	result := []Entry{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	if len(result) == 0 {
		return nil, ErrNoDefinitionsFound
	}

	return result, nil
}

func GetWordEntry(ctx context.Context, language, word string) (*Entry, error) {
	entries, err := fetchWordEntries(ctx, language, word)
	if err != nil {
		return nil, err
	}

	return &entries[0], nil
}
