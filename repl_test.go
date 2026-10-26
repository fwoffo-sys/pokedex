package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Goodbye  world  ",
			expected: []string{"goodbye", "world"},
		},
		{
			input:    "  helloWorld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "  hello  world\n  and goodbye world",
			expected: []string{"hello", "world", "and", "goodbye", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length Error: actual was '%v', expected was '%v'", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Correct Error: cleaned word %v not equal to expected word %v from %v", actual, c.expected, c.input)
			}
		}
	}
}