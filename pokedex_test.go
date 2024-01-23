package main

import "testing"

func TestParseInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := parseInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Lengths are not equal: %v and %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			expWord := cs.expected[i]
			actWord := actual[i]
			if expWord != actWord {
				t.Errorf("Words do not match. Index: %v. Expected: %v. Acutal: %v", i, expWord, actWord)
				break
			}
		}
	}
}