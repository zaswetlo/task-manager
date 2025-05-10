package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"task-manager/internal/model"
)

var tasks = []model.Task{}
var nextID = 1

const storageFile = "tasks.json"

func init() {
	loadTasks()
}

func loadTasks() {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		return
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		panic(err)
	}

	// Find the highest ID to set nextID
	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	// Ensure the directory exists
	dir := filepath.Dir(storageFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0644)
}

func GetAllTasks() []model.Task {
	return tasks
}

func AddTask(title string) model.Task {
	task := model.Task{ID: nextID, Title: title, Description: "", Done: false}
	nextID++
	tasks = append(tasks, task)

	if err := saveTasks(); err != nil {
		panic(err)
	}

	return task
}

func UpdateTask(id int, done bool) (model.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = done

			if err := saveTasks(); err != nil {
				panic(err)
			}

			return tasks[i], nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func UpdateTaskDetails(id int, title, description string) (model.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Description = description

			if err := saveTasks(); err != nil {
				panic(err)
			}

			return tasks[i], nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks()
		}
	}
	return errors.New("task not found")
}

func GetTaskProgress() float64 {
	if len(tasks) == 0 {
		return 0
	}

	completed := 0
	for _, task := range tasks {
		if task.Done {
			completed++
		}
	}

	return float64(completed) / float64(len(tasks)) * 100
}
