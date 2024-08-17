package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreateAt    time.Time
	CompletedAt time.Time
}

type List []item

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

	list[len(list)-1].Done = true
	list[len(list)-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(index int) error {
	list := *l
	if len(list) == 0 || index > len(list) {
		return errors.New("invalid index")
	}

	*l = append(list[:index-1], list[index:]...)

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
