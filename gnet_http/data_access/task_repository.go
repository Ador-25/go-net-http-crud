package data_access

import (
	"encoding/json"
	"fmt"
	"net_http/models"
	"time"
)

var tasks []models.Task

func init() {
	tasks = []models.Task{
		{Id: 1, Type: "dotnet", StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Description: "Task 1", IsCompleted: false},
		{Id: 2, Type: "golang", StartTime: time.Now(), EndTime: time.Now().Add(time.Hour), Description: "Task 2", IsCompleted: true},
	}
}
func GetAllTasks() ([]byte, error) {
	data, err := json.Marshal(tasks)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetTask(id int) ([]byte, error) {
	data, err := json.Marshal(tasks[id-1])
	if err != nil {
		return nil, err
	}
	return data, nil
}

func AddTask(task models.Task) {
	task.Id = len(tasks) + 1
	tasks = append(tasks, task)
}

func EditTask(id int, task models.Task) {
	current := &tasks[id-1]
	current.Description = task.Description
	current.EndTime = task.EndTime
	current.StartTime = task.StartTime
	current.Type = task.Type
}

func RemoveTask(id int) {
	index := id - 1
	tasks = append(tasks[0:index], tasks[index+1:]...)
	fmt.Println(tasks)
}
