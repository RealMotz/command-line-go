package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := List{}

	task := "New Task"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %s, got %s", task, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := List{}

	task := "New Task"
	l.Add(task)
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := List{}

	l.Add("New Task")
	l.Add("New Task 2")
	l.Delete(1)

	if len(l) != 1 {
		t.Errorf("New Task should be deleted")
	}
}

func TestSaveGet(t *testing.T) {
	file, err := os.CreateTemp("", "")
	if err != nil {
		t.Errorf("Error creating temp file %v: ", err)
	}
	defer os.Remove(file.Name())

	l := List{}
	l.Add("New Task")
	err = l.Save(file.Name())
	if err != nil {
		t.Errorf("Error saving file %v: ", err)
	}

	l2 := List{}
	err = l2.Get(file.Name())
	if err != nil {
		t.Errorf("Error getting file %v: ", err)
	}

	if l[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l[0].Task, l2[0].Task)
	}
}
