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
func main() {
	fmt.Println("Debug: ", os.Args)
	switch {
	case len(os.Args) < 2:
		fmt.Println("usage")
		return
	case os.Args[1] == "add":
		fmt.Println("Add")
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
		fmt.Println("Update")
		if (len(os.Args) > 3){
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				
				for i, t := range tasks {
					if t.Id == id {
						tasks[i].Description = os.Args[3]
						tasks[i].UpdatedAt = time.Now()
						saveTasks(tasks)
						fmt.Printf("Task #%s was successfully updated!\n", os.Args[2])
						break
					}
				}
			}
			
		}
	case os.Args[1] == "delete":
		if (len(os.Args) > 2){
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				
				for i, t := range tasks {
					if t.Id == id {
						tasks = append(tasks[:i], tasks[i+1:]...)
						saveTasks(tasks)
						fmt.Printf("Task #%s was deleted successfully \n", os.Args[2])
						break
					}
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
		if (len(os.Args) > 2){
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				
				for i, t := range tasks {
					if t.Id == id {
						tasks[i].Status = "in progress"
						tasks[i].UpdatedAt = time.Now()
						saveTasks(tasks)
						fmt.Printf("Task #%s was successfully updated!\n", os.Args[2])
						break
					}
				}
			}
			
		}
	case os.Args[1] == "mark-done":
		if (len(os.Args) > 2){
			tasks := loadTasks()
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
			} else {
				
				for i, t := range tasks {
					if t.Id == id {
						tasks[i].Status = "done"
						tasks[i].UpdatedAt = time.Now()
						saveTasks(tasks)
						fmt.Printf("Task #%s was successfully updated!\n", os.Args[2])
						break
					}
				}
			}
			
		}
    }
	
	
		
}


		
		
	

