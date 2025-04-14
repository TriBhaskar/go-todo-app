package todo

import (
	"encoding/json"
	"os"
)

var todoFile string

func SetTodoFile(path string) {
	todoFile = path
}

// SaveToFile saves the todo list to the specified file
func SaveToFile() error {
	data, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

// LoadFromFile loads the todo list from the specified file
func LoadFromFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		todoList = TodoList{Todos: []Todo{}}
		return SaveToFile()
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &todoList)
}