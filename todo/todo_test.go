package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.TodoList{}

	task := "Work on leetcode"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %s, got %s", task, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.TodoList{}

	task := "Work on leetcode"
	l.Add(task)

	if l[0].Done {
		t.Errorf("Task shouldn't completed without calling complete")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be completed after calling complete")
	}
}

func TestDeleteMethod(t *testing.T) {
	l := todo.TodoList{}

	task1 := "New task"
	task2 := "New task"

	l.Add(task1)
	l.Add(task2)

	l.Delete(0)

	if len(l) != 1 {
		t.Errorf("The length of list should be one after deleting a task from a task list of length 2")
	}

	if l[0].Task != task2 {
		t.Errorf("The task2 should be moved to task1 places after deleting the task1")
	}
}

func TestGetSave(t *testing.T) {
	currentList := todo.TodoList{}

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	for _, task := range tasks {
		currentList.Add(task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Cannot create temporary file")
	}

	err = currentList.Save(tf.Name())

	if err != nil {
		t.Errorf("Cannot save to temporary file")
	}

	loadedList := todo.TodoList{}

	err = loadedList.Get(tf.Name())

	if err != nil {
		t.Errorf("Cannot load from saved file")
	}

	for i := range currentList {
		if currentList[i].Task != loadedList[i].Task {
			t.Errorf("current list and loaded list dosn't get matched")
		}
	}
}
