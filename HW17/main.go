package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	tasks   = []Task{}
	nextID  = 1
	mu      sync.Mutex
	rdb     *redis.Client
	ctx     = context.Background()
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func getTasks(c echo.Context) error {
	cacheKey := "tasks"
	cachedTasks, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		mu.Lock()
		defer mu.Unlock()
		tasksJSON, err := json.Marshal(tasks)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error encoding tasks"})
		}
		err = rdb.Set(ctx, cacheKey, tasksJSON, 5*time.Minute).Err()
		if err != nil {
			c.Logger().Error("Error saving to cache:", err)
		}
		return c.JSON(http.StatusOK, tasks)
	} else if err != nil {
		c.Logger().Error("Error fetching from cache:", err)
		return c.JSON(http.StatusOK, tasks)
	}

	var cachedTasksList []Task
	err = json.Unmarshal([]byte(cachedTasks), &cachedTasksList)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error decoding cache"})
	}
	return c.JSON(http.StatusOK, cachedTasksList)
}

func addTask(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task data"})
	}
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	err := rdb.Del(ctx, "tasks").Err()
	if err != nil {
		c.Logger().Error("Error clearing cache:", err)
	}

	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task ID"})
	}
	mu.Lock()
	defer mu.Unlock()
	var updatedTask Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task data"})
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed
			err := rdb.Del(ctx, "tasks").Err()
			if err != nil {
				c.Logger().Error("Error clearing cache:", err)
			}
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task ID"})
	}
	mu.Lock()
	defer mu.Unlock()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			err := rdb.Del(ctx, "tasks").Err()
			if err != nil {
				c.Logger().Error("Error clearing cache:", err)
			}
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", addTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}
