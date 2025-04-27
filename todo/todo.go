package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreateAt    time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) String() string {
	return l.ToString(false)
}

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreateAt:    time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(index int) error {
	list := *l
	if len(list) == 0 || index > len(list) {
		return errors.New("invalid index")
	}

	list[index-1].Done = true
	list[index-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(index int) error {
	list := *l
	if len(list) == 0 || index > len(list) {
		return errors.New("invalid index")
	}

	*l = slices.Delete(list, index-1, index)

	return nil
}

func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, l)
}

func (l *List) ToString(verbose bool) string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		suffix := ""
		if t.Done {
			prefix = "X"
			suffix = t.CompletedAt.String()
		}

		if verbose {
			formatted += fmt.Sprintf("[%s] %d: %s %s (%s)\n", prefix, k+1, t.Task, t.CreateAt.String(), suffix)
		} else {
			// Adjust the item number k to print numbers starting from 1 instead of 0
			formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
		}
	}

	return formatted
}
