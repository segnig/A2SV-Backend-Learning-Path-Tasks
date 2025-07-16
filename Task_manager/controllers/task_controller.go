package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	tasks := data.GetAllTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	if task, found := data.GetTaskById(id); found {
		ctx.JSON(http.StatusOK, task)
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func CreateTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, OK := data.AddNewTask(newTask)

	if !OK {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task already exists"})
		return
	}
	ctx.JSON(http.StatusCreated, createdTask)
}

func DeleTeTask(ctx *gin.Context) {
	id := ctx.Param("id")

	_, OK := data.DeleteTaskById(id)

	if !OK {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, OK := data.UpdateTaskbyId(id, updatedTask)

	if !OK {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
