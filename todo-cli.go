package main

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
)

type Task struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func loadTasks() []Task {
	var tasks []Task
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println(err)
		return tasks
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println(err)
		return tasks
	}
	return tasks
		
} 
func main() {
	fmt.Println("Debug: ", os.Args)
	switch {
	case len(os.Args) < 2:
		fmt.Println("usage")
		return
	case os.Args[1] == "add":
		fmt.Println("Add")
		tasks := loadTasks()
		fmt.Println("debug tasks: ", tasks)
	case os.Args[1] == "update":
		fmt.Println("Update")
	case os.Args[1] == "delete":
		fmt.Println("delete")
	}
}
