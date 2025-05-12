package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
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

	if err := run(*filename, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(file string, out io.Writer) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	html := parseContent(content)

	temp, err := os.CreateTemp(filepath.Dir(file), "mdp*.html")
	if err != nil {
		fmt.Println("Error creating a temp file")
		return err
	}
	defer temp.Close()

	fmt.Fprintln(out, temp.Name())

	return os.WriteFile(temp.Name(), html, 0644)
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
