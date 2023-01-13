package reversoapi_test

import (
	"context"
	"testing"

	"github.com/rkfcccccc/english_words/services/dictionary/pkg/reversoapi"
	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	client := reversoapi.NewClient()

	result, err := client.GetTranslation(context.Background(), "test", "eng", "rus")

	assert.Nil(t, err)
	assert.Equal(t, result.ContextResults.Results[0].Translation, "тест")
}
