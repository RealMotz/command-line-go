package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4

	actual := count(b, false, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nword4\nword5")
	expected := 3

	actual := count(b, true, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("0123456789")
	expected := 10

	actual := count(b, false, true)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}

func TestCountWordStdinError(t *testing.T) {
	tests := []struct {
		name     string
		input    *strings.Reader
		lFlag    bool
		bFlag    bool
		expected int
	}{
		{
			name:     "buffer max scan token size reached",
			input:    strings.NewReader(strings.Repeat("a", bufio.MaxScanTokenSize+1)),
			lFlag:    false,
			bFlag:    false,
			expected: 0,
		},
		{
			name:     "-l and -b flags used together",
			input:    strings.NewReader("test string"),
			lFlag:    true,
			bFlag:    true,
			expected: 0,
		},
	}

	for _, tc := range tests {
		actual := count(tc.input, tc.lFlag, tc.bFlag)
		if actual != tc.expected {
			t.Errorf("%s: Expected %d, got %d instead. \n", tc.name, tc.expected, actual)
		}
	}
}
