package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}

	got := parseContent(content)
	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, got) {
		t.Logf("golden: \n%s\n", expected)
		t.Logf("got: \n%s\n", got)
		t.Error("Result content does not match golden file")
	}
}

func TestRun(t *testing.T) {
	var mockStdOut bytes.Buffer
	got := run(inputFile, &mockStdOut)
	if got != nil {
		t.Logf("error: %v", got)
		t.Fatal("Error converting to html")
	}

	resultFile := strings.TrimSpace(mockStdOut.String())

	actual, err := os.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expected, actual) {
		t.Logf("golden: \n%s\n", expected)
		t.Logf("got: \n%s\n", got)
		t.Error("Result content does not match golden file")
	}

	os.Remove(resultFile)
}
