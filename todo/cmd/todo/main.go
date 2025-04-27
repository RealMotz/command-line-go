package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/RealMotz/command-line-go/todo"
)

// Default filename
var filename = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "todo CLI. developed for the Pragmatic Bookshelf\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	if os.Getenv("TODO_FILENAME") != "" {
		filename = os.Getenv("TODO_FILENAME")
	}

	task := flag.Bool("a", false, "Create a new task")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Complete a task")
	del := flag.Int("d", -1, "Delete a task")
	verbose := flag.Bool("v", false, "Verbose")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
			fmt.Print(l.ToString(*verbose))
	case *task:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(task)

		if err := l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		l.Complete(*complete)

		if err := l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del >= 0:
		err := l.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		joinedArgs := strings.Join(args, " ")
		return joinedArgs, nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be empty")
	}

	return s.Text(), nil
}
