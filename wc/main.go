package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines, countBytes, wordCount, countRunes bool) int {
	if countLines && countBytes {
		fmt.Fprintln(os.Stderr, "invalid flags")
		return 0
	}

	// Used to read text from a reader
	scanner := bufio.NewScanner(r)

	if wordCount {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	if countRunes {
		scanner.Split(bufio.ScanRunes)
	}

	// word count
	wc := 0

	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		return 0
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	words := flag.Bool("w", false, "Count words")
	runes := flag.Bool("r", false, "Count runes")

	// Parsing the flags provided by the user
	flag.Parse()

	filename := os.Args[len(os.Args)-1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(count(os.Stdin, *lines, *bytes, *words, *runes))
		return
	}
	defer f.Close()

	// Calling the count function to count the number of words
	// Received from the Standard Input and printing it out
	// fmt.Println(count(os.Stdin, *lines, *bytes, *words))
	if !(*lines || *bytes || *words || *runes) {
		b := count(f, false, true, false, false)
		f.Seek(0, io.SeekStart)
		w := count(f, false, false, true, false)
		f.Seek(0, io.SeekStart)
		l := count(f, true, false, false, false)
		fmt.Printf("%d %d %d %s\n", l, w, b, filename)
	} else {
		fmt.Println(count(f, *lines, *bytes, *words, *runes))
	}
}
