package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed getting tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := data.GetTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := data.AddNewTask(newTask)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task already exists"})
		return
	}
	ctx.JSON(http.StatusCreated, createdTask)
}

func DeleTeTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := data.DeleteTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
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

	err := data.UpdateTaskbyId(id, updatedTask)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}
