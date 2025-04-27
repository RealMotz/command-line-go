package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<head>
<meta http-equiv="content-type" content="text/html; charset=utf-8">
<title>Markdown Preview Tool</title>
</head>
<body>`
	footer = `</body>
</html>`
)

func main() {
	filename := flag.String("f", "", "Parse file")
	flag.Parse()

	if len(*filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	html := parseContent(content)

	output := fmt.Sprintf("%s.html", filepath.Base(filename))

	return os.WriteFile(output, html, 0644)
}

func parseContent(input []byte) []byte {
	unsafe := blackfriday.Run(input, blackfriday.WithNoExtensions())
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(html)
	buffer.WriteString(footer)

	return buffer.Bytes()
}
