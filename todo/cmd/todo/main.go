package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RealMotz/command-line-go/todo"
)

// harcoding the name
const filename = ".todo.json"

func main() {
	task := flag.String("task", "", "Create a new task")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Complete a task")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	case *task != "":
		l.Add(*task)

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
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
