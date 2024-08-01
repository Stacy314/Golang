/*Additional materials:
https://redis.io/commands/;
https://redis.io/docs/latest/develop/connect/clients/go/

Add the Redis cache to the REST API implementation in the HW14 (Layered REST) ​​or HW10 (REST) ​​task:
• we receive a request from the client;
• check whether there is data in the cache, if there is — return it, if not — contact the database;
• return the response from the database, while storing it in redis so that the database is not accessed next time.*/


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

func fetchFromCache(cacheKey string) ([]Task, error) {
	cachedTasks, err := rdb.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var cachedTasksList []Task
	err = json.Unmarshal([]byte(cachedTasks), &cachedTasksList)
	if err != nil {
		return nil, err
	}
	return cachedTasksList, nil
}

func saveToCache(cacheKey string, tasks []Task) error {
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return rdb.Set(ctx, cacheKey, tasksJSON, 5*time.Minute).Err()
}

func clearCache(cacheKey string) error {
	return rdb.Del(ctx, cacheKey).Err()
}

func getTasks(c echo.Context) error {
	cacheKey := "tasks"
	cachedTasks, err := fetchFromCache(cacheKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error fetching from cache"})
	}
	if cachedTasks == nil {
		mu.Lock()
		defer mu.Unlock()
		return handleCacheMiss(c, cacheKey)
	}
	return c.JSON(http.StatusOK, cachedTasks)
}

func handleCacheMiss(c echo.Context, cacheKey string) error {
	err := saveToCache(cacheKey, tasks)
	if err != nil {
		c.Logger().Error("Error saving to cache:", err)
	}
	return c.JSON(http.StatusOK, tasks)
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

	if err := clearCache("tasks"); err != nil {
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
			if err := clearCache("tasks"); err != nil {
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
			if err := clearCache("tasks"); err != nil {
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

