package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"todo/cmd/todo"
)

func main() {
	createFlg := flag.Bool("c", false, "create a new todo")
	_ = flag.Bool("d", false, "deleate a todo")
	flag.Parse()

	var todo todo.Todo
	switch {
	case len(os.Args) == 1:
		for _, item := range todo.List {
			fmt.Println("printing??")
			fmt.Println(item.Content)
		}
	default:
		fmt.Println(strings.Join(os.Args[1:], " "))
	}

	if *createFlg {
		todo.Add(strings.Join(flag.Args(), " "))
	}
}
