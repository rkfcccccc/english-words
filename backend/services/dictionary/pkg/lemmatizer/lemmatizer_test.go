package lemmatizer

import (
	"testing"
)

func TestLemma(t *testing.T) {
	lemmatizer := New("en")

	testCases := []struct{ Input, Output string }{
		{"attacked", "attack"},
		{"assessing", "assess"},
		{"asserts", "assert"},
	}

	for _, testCase := range testCases {
		if s := lemmatizer.Lemma(testCase.Input); s != testCase.Output {
			t.Fatalf("%s: expected %s got %s", testCase.Input, testCase.Output, s)
		}
	}
}
