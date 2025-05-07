package storage

import "task-manager/internal/model"

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
