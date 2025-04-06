package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type List []TodoItem

type TodoItem struct {
	Id          uuid.UUID `json:"id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type Todo struct {
	List List
}

func (t *Todo) Add(content string) TodoItem {
	todo := TodoItem{
		Id:        uuid.New(),
		Content:   content,
		CreatedAt: time.Time{},
	}

	file, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(todo)
	if err != nil {
		fmt.Println("Error encoding todo item:", err)
	}

	t.List = append(t.List, todo)
	return todo
}

func (t *Todo) Delete(id uuid.UUID) {
	fmt.Println("deleating...")
}

func (t *Todo) Save(id uuid.UUID) {
	fmt.Println("saving...")
}

func (t *Todo) Get(id uuid.UUID) TodoItem {
	fmt.Println("getting...")
	return TodoItem{}
}

func (t *Todo) Complete(id uuid.UUID) {
	fmt.Println("Completing...")
}
