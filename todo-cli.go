package main

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	"strconv"
)

type Task struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	StatusToDo = "todo"
	StatusInProgress = "in-progress"
	StatusDone = "done"
)

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

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("saveTasks Error: ", err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("saveTasks Error: ", err)
		return
	}

}

func findTaskIndex (tasks []Task, id int) int {
	for i,t := range tasks {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func changeStatus (idString string, status string) {
	tasks := loadTasks()
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	} else {
		idx := findTaskIndex(tasks, id)
		if idx >= 0 {
			tasks[idx].Status = status
			tasks[idx].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task #%s was successfully updated!\n", os.Args[2])
			} else {
				fmt.Println("Task not found")
			}
		}
}
func main() {
	switch {
	case len(os.Args) < 2:
		fmt.Println("Wrong arguments, read README.md to use it correctly")
		return
	case os.Args[1] == "add":
		if len(os.Args) >= 3 {
			tasks := loadTasks()
			maxID := 0
			for _, t := range tasks {
    			if t.Id > maxID {
       				 maxID = t.Id
				}
			}
			now := time.Now()
			newTask := Task {
				Id: maxID + 1,
				Description: os.Args[2],
				Status: "todo",
				CreatedAt: now,
				UpdatedAt: now,
			}
			tasks = append(tasks, newTask)
			saveTasks(tasks)
		}
	case os.Args[1] == "update":
		if  len(os.Args) > 3 {
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				idx := findTaskIndex(tasks, id)
				if idx >= 0 {
					tasks[idx].Description = os.Args[3]
					tasks[idx].UpdatedAt = time.Now()
					saveTasks(tasks)
					fmt.Printf("Task #%s was successfully updated!\n", os.Args[2])
				} else {
					fmt.Println("Task not found")
				}
			}
		}
	case os.Args[1] == "delete":
		if  len(os.Args) > 2 {
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				idx := findTaskIndex(tasks, id) 
				if idx >= 0 {
					tasks = append(tasks[:idx], tasks[idx+1:]...)
					saveTasks(tasks)
					fmt.Printf("Task #%s was deleted successfully \n", os.Args[2])
				} else {
					fmt.Println("Task not found")
				}
			}
		}
	case os.Args[1] == "list":
		tasks := loadTasks()
		var result []Task
		if (len(os.Args) > 2){
			for _, t := range tasks {
				if t.Status == os.Args[2] {
					result = append(result, t)
				}
			}
		} else {
			result = tasks
		}
		if len(result) == 0 {
			fmt.Println("No tasks")
		} else {
			for _, t := range result {
				fmt.Printf("[#%d] [%s] %s \n", t.Id, t.Status, t.Description)
			}
		}
	case os.Args[1] == "mark-in-progress":
		if len(os.Args) > 2 {
			changeStatus(os.Args[2], StatusInProgress)
		}
	case os.Args[1] == "mark-done":
		if len(os.Args) > 2 {
			changeStatus(os.Args[2], StatusDone)
		}
    }
}