package srt

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var (
	filenameRegex = regexp.MustCompile(`filename="(.*)"`)
	htmlTagRegex  = regexp.MustCompile("<[^>]*>")
)

func removeHtmlTags(s string) string {
	return htmlTagRegex.ReplaceAllString(s, "")
}

func containsLetters(s string) bool {
	return 'a' <= s[0] && s[0] <= 'z' && len(s) > 2
}

func extractWords(content string) []string {
	parts := strings.Split(content, "\n\n")

	seen := map[string]struct{}{}
	for _, part := range parts {
		part = removeHtmlTags(part)
		lines := strings.Split(part, "\n")
		sentences := lines[2:]

		for _, sentence := range sentences {
			sentence = replacer.Replace(sentence)

			words := strings.Fields(sentence)
			for _, word := range words {
				word = strings.ToLower(word)

				if word[len(word)-1] == '\'' {
					word = word[:len(word)-1] + "g"
				}

				if containsLetters(word) {
					seen[word] = struct{}{}
				}
			}
		}
	}

	words := make([]string, len(seen))

	i := 0
	for word := range seen {
		words[i] = word
		i++
	}

	return words
}

func unzipContents(body []byte) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("zip.NewReader: %v", err)
	}

	if len(zipReader.File) != 1 {
		return nil, fmt.Errorf("zip contains more or less than 1 file: %d", len(zipReader.File))
	}

	zippedFile := zipReader.File[0]
	if !strings.HasSuffix(zippedFile.Name, ".srt") {
		return nil, fmt.Errorf("zip file contains non .srt file: %s", zippedFile.Name)
	}

	f, err := zippedFile.Open()
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll from zipped file: %v", err)
	}

	defer f.Close()
	body, err = io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll from zipped file: %v", err)
	}

	return body, nil
}

func GetWordsFromURL(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}

	disposition := response.Header.Get("Content-Disposition")
	matches := filenameRegex.FindStringSubmatch(disposition)

	if len(matches) != 2 {
		return nil, fmt.Errorf("could not parse Content-Disposition: %s", disposition)
	}

	filename := matches[1]
	switch {
	case strings.HasSuffix(filename, ".zip"):
		if body, err = unzipContents(body); err != nil {
			return nil, err
		}
	case strings.HasSuffix(filename, ".srt"):
	default:
		return nil, fmt.Errorf("unknown extension of file: %s", filename)
	}

	return extractWords(string(body)), nil
}
