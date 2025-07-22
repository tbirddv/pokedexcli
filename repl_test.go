package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"  Hello World  ", []string{"hello", "world"}},
		{"Go is great", []string{"go", "is", "great"}},
		{"  Multiple   spaces   ", []string{"multiple", "spaces"}},
		{"Mixed CASE Input", []string{"mixed", "case", "input"}},
	}

	for _, test := range tests {
		result := cleanInput(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("Expected %d elements, got %d for input '%s'", len(test.expected), len(result), test.input)
			continue
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Expected '%s', got '%s' for input '%s'", test.expected[i], v, test.input)
			}
		}
	}
}
