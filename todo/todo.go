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
	CreatedAt   time.Time
	Done        bool
	CompletedAt time.Time
}

type TodoList []item

func (list *TodoList) Complete(id int) error {
	if id <= 0 || id > len(*list) {
		return fmt.Errorf("%d is not a valid id", id)
	}

	(*list)[id - 1].Done = true
	(*list)[id - 1].CompletedAt = time.Now()

	return nil
}

func (list *TodoList) Add(task string) {
	t := item{
		Task:        task,
		CreatedAt:   time.Now(),
		Done:        false,
		CompletedAt: time.Time{},
	}

	*list = append(*list, t)
}

func (list *TodoList) Delete(id int) error {
	if id <= 0 || id > len(*list) {
		return fmt.Errorf("%d is not a valid id", id)
	}

	*list = slices.Delete((*list), id - 1, id)

	return nil

}

func (list *TodoList) Save(filename string) error {
	content, err := json.Marshal(list)

	if err != nil {
		return err
	}

	os.WriteFile(filename, content, 0644)

	return nil
}

func (list *TodoList) Get(filename string) error {
	filecontent, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	return json.Unmarshal(filecontent, list)

}


func (l *TodoList) String() string{

	formatted := ""

	for i, item := range *l {
		header := "[ ]"

		if item.Done {
			header = "[x]"
		}

		taskString := fmt.Sprintf("%d %s %s\n", i + 1, header, item.Task)

		formatted += taskString
	}

	return formatted
}
