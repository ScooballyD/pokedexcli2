package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "   ",
			expected: "",
		},
		{
			input:    "  hello  ",
			expected: "hello",
		},
		{
			input:    "  HeLlo  ",
			expected: "hello",
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v vs %v'", actual, c.expected)
			continue
		}
		if actual != c.expected {
			t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
		}
	}
}
