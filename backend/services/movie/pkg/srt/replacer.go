package srt

import "strings"

var replaceList = []string{
	",", " ",
	".", " ",
	"?", " ",
	"!", " ",
	"'s", "",
	"'d", " would",
	"'ll", " will",
	"'re", " are",
	"'ve", " have",
	"'m", "",
	":", "",
	`"`, "",

	"[", " ", "]", " ",
	"(", " ", ")", " ",
	"{", " ", "}", " ",

	"won't", "will not",
	"n't", "  not",
}

var replacer = strings.NewReplacer(replaceList...)
