package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := "4"
	res := wordCount(b, false, false, true, false)

	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func TestLineCount(t *testing.T) {
	b := bytes.NewBufferString("word1\nword2\nword3\nword4")
	exp := "4"
	res := wordCount(b, false, false, false, true)

	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func TestCharCount(t *testing.T) {
	b := bytes.NewBufferString("1234567890\n")
	exp := "11"
	res := wordCount(b, false, true, false, false)

	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}

func TestByteCount(t *testing.T) {
	b := bytes.NewBufferString("1234567890\n")
	exp := "11"
	res := wordCount(b, true, false, false, false)

	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}
