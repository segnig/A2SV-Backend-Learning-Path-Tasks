package data

import (
	"task-manager/models"
	"time"
)

var tasks = []models.Task{
	{
		ID:          "1",
		Title:       "Finish project proposal",
		Description: "Write and submit the proposal to the client",
		DueDate:     time.Date(2025, 7, 20, 17, 0, 0, 0, time.UTC),
		Status:      "Pending",
	},
	{
		ID:          "2",
		Title:       "Team meeting",
		Description: "Weekly sync with the development team",
		DueDate:     time.Date(2025, 7, 18, 10, 0, 0, 0, time.UTC),
		Status:      "In Progress",
	},
	{
		ID:          "3",
		Title:       "Update project documentation",
		Description: "Revise API reference and architecture diagrams",
		DueDate:     time.Date(2025, 7, 25, 15, 0, 0, 0, time.UTC),
		Status:      "Pending",
	},
	{
		ID:          "4",
		Title:       "Deploy new version",
		Description: "Deploy v1.2.3 to staging environment",
		DueDate:     time.Date(2025, 7, 22, 13, 0, 0, 0, time.UTC),
		Status:      "Completed",
	},
	{
		ID:          "5",
		Title:       "Code review",
		Description: "Review PRs assigned in the sprint board",
		DueDate:     time.Date(2025, 7, 19, 9, 30, 0, 0, time.UTC),
		Status:      "In Progress",
	},
}

// Get all tasks
func GetAllTasks() []models.Task {
	return tasks
}

// Get specific task by its id
func GetTaskById(id string) (models.Task, bool) {

	for _, task := range tasks {
		if task.ID == id {
			return task, true
		}
	}

	return models.Task{}, false
}

// update specific task by its id
func UpdateTaskbyId(id string, updatedTask models.Task) (models.Task, bool) {

	for idx, task := range tasks {
		if task.ID == id {
			if task.Title != "" {
				tasks[idx].Title = updatedTask.Title
			}

			if task.Description != "" {
				tasks[idx].Description = updatedTask.Description
			}
			if task.Status != "" {
				tasks[idx].Status = updatedTask.Status
			}

			return tasks[idx], true
		}
	}

	return models.Task{}, false
}

// Delete task by its id if exists
func DeleteTaskById(id string) (models.Task, bool) {
	deletedTask := models.Task{}

	for idx, task := range tasks {
		if task.ID == id {
			deletedTask = task
			tasks = append(tasks[:idx], tasks[idx+1:]...)

			return deletedTask, true
		}
	}

	return deletedTask, false
}

// Add new task if not exists
func AddNewTask(newTask models.Task) (models.Task, bool) {
	for _, task := range tasks {
		if task.ID == newTask.ID {
			return models.Task{}, false
		}
	}
	tasks = append(tasks, newTask)
	return newTask, true
}
