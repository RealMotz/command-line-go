package todo_test

import (
	"testing"
	"todo/cmd/todo"
)

func TestAdd(t *testing.T) {
	l := todo.Todo{}

	taskName := "new task"
	l.Add(taskName)

	if l.List[0].Content != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l.List[0].Content)
	}
}


