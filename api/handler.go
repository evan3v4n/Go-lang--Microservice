// Path: project_root/api/handlers/task_handler.go

package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"Go_lang_Microservice/api/models"
	"Go_lang_Microservice/db"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := db.DB.Create(&task).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Cache the new task
	taskJSON, _ := json.Marshal(task)
	db.RedisClient.Set(context.Background(), "task:"+task.ID, taskJSON, 1*time.Hour)

	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := db.DB.Find(&tasks).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	// Try to get task from cache
	cachedTask, err := db.RedisClient.Get(context.Background(), "task:"+id).Result()
	if err == nil {
		var task models.Task
		json.Unmarshal([]byte(cachedTask), &task)
		c.JSON(http.StatusOK, task)
		return
	}

	// If not in cache, get from database
	var task models.Task
	if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Cache the task
	taskJSON, _ := json.Marshal(task)
	db.RedisClient.Set(context.Background(), "task:"+id, taskJSON, 1*time.Hour)

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed
	task.UpdatedAt = time.Now()

	if err := db.DB.Save(&task).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	// Update cache
	taskJSON, _ := json.Marshal(task)
	db.RedisClient.Set(context.Background(), "task:"+id, taskJSON, 1*time.Hour)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := db.DB.First(&task, "id = ?", id).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := db.DB.Delete(&task).Error; err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	// Remove from cache
	db.RedisClient.Del(context.Background(), "task:"+id)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
