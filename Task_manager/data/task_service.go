package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	taskCollection mongo.Collection
	cxt            = context.Background()
)

func init() {
	mongoURL := os.Getenv("MONGO_URL")

	if mongoURL == "" {
		mongoURL = "mongodb://localhost:27017"
	}

	var err error

	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err = mongo.Connect(cxt, clientOptions)

	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	taskCollection = *client.Database("task-manager").Collection("tasks")
}

// Get all tasks
func GetAllTasks() ([]models.Task, error) {

	cursor, err := taskCollection.Find(cxt, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(cxt)

	var tasks []models.Task

	if err = cursor.All(cxt, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Get specific task by its id
func GetTaskById(id string) (*models.Task, error) {

	var task models.Task
	err := taskCollection.FindOne(cxt, bson.M{"id": id}).Decode(&task)
	return &task, err
}

// update specific task by its id
func UpdateTaskbyId(id string, updatedTask models.Task) error {

	result, err := taskCollection.UpdateOne(
		cxt,
		bson.M{"id": id},
		bson.M{"$set": updatedTask},
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no task found with id '%s'", id)
	}

	return err
}

// Delete task by its id if exists
func DeleteTaskById(id string) error {
	result, err := taskCollection.DeleteOne(
		cxt,
		bson.M{"id": id},
	)

	if result.DeletedCount == 0 {
		return fmt.Errorf("task id '%v' not found", id)
	}

	return err
}

// Add new task if not exists
func AddNewTask(newTask models.Task) (*models.Task, error) {
	count, err := taskCollection.CountDocuments(cxt, bson.M{"id": newTask.ID})

	if err != nil {
		return nil, fmt.Errorf("failed to check task ID uniqueness %v", err)
	}

	if count > 0 {
		return nil, fmt.Errorf("task ID '%s' already exists", newTask.ID)
	}

	_, err = taskCollection.InsertOne(cxt, newTask)

	return &newTask, err
}
