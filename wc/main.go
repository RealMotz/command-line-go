package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	// Used to read text from a reader
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
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
	// Defining a bollean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// Received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines))
}
