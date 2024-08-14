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

	actual := count(b, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nword4\nword5")
	expected := 3

	actual := count(b, true)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}

func TestCountWordStdinError(t *testing.T) {
	overflow := strings.Repeat("a", bufio.MaxScanTokenSize+1)
	r := strings.NewReader(overflow)
	expected := 0

	actual := count(r, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead. \n", expected, actual)
	}
}
