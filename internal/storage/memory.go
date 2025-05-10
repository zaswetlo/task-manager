package storage

import (
	"errors"
	"task-manager/internal/model"
)

var tasks = []model.Task{}
var nextID = 1

func GetAllTasks() []model.Task {
	return tasks
}

func AddTask(title string) model.Task {
	task := model.Task{ID: nextID, Title: title, Done: false}
	nextID++
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, done bool) (model.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = done
			return tasks[i], nil
		}
	}
	return model.Task{}, errors.New("task not found")
}
