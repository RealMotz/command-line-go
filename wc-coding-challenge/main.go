package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	countBytes := flag.Bool("c", false, "count number of bytes")
	countWords := flag.Bool("w", false, "count number of words")
	countChar := flag.Bool("m", false, "count number of chars")
	countLines := flag.Bool("l", false, "count number of lines")

	flag.Parse()

	args := flag.Args()
	var reader io.Reader

	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}

	fmt.Println(wordCount(reader, *countBytes, *countChar, *countWords, *countLines))
}

func wordCount(f io.Reader, totalBytes, totalChars, totalWords, totalLines bool) string {
	reader := bufio.NewReader(f)

	switch {
	case totalBytes:
		return countBytes(reader)
	case totalWords:
		return countLines(reader, func(line string) int { return len(strings.Fields(line)) })
	case totalChars:
		return countLines(reader, func(line string) int { return len(line) })
	case totalLines:
		return countLines(reader, func(string) int { return 1 })
	default:
		return defaultCount(reader)
	}
}

func defaultCount(r *bufio.Reader) string {
	var lines, words, chars int
	for {
		line, err := r.ReadString('\n')
		chars += len(line)
		words += len(strings.Fields(line))
		lines++
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return fmt.Sprintf("%d %d %d", lines, words, chars)
}

func countLines(r *bufio.Reader, fn func(string) int) string {
	var count int
	for {
		line, err := r.ReadString('\n')
		count += fn(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return fmt.Sprintf("%d", count)
}

func countBytes(r *bufio.Reader) string {
	var count int
	for {
		line, err := r.ReadBytes('\n')
		count += len(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return fmt.Sprintf("%d", count)
}
