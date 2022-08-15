package lemmatizer

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed lemmatization-*.txt
var fs embed.FS

type Lemmatizer struct {
	m map[string]string
}

func New(language string) *Lemmatizer {
	lem := &Lemmatizer{map[string]string{}}

	if err := lem.init(language); err != nil {
		log.Fatalf("lemmatizer failed init with language %s: %v", language, err)
	}

	return lem
}

func (l *Lemmatizer) init(language string) error {
	bytes, err := fs.ReadFile(fmt.Sprintf("lemmatization-%s.txt", language))
	if err != nil {
		return fmt.Errorf("os.ReadFile: %v", err)
	}

	lines := strings.Split(string(bytes), "\n")
	for i, line := range lines {
		words := strings.Split(strings.TrimSpace(line), "\t")

		if len(words) != 2 {
			return fmt.Errorf("failed parsing words at line %d", i+1)
		}

		l.m[words[1]] = words[0]
	}

	return nil
}

func (l *Lemmatizer) Lemma(word string) string {
	lemma, ok := l.m[strings.ToLower(word)]

	if !ok {
		return word
	}

	return lemma
}
