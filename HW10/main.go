/*Create a server with a REST API to view a list of rights.
The user must be able to:
* view the list of tasks
* add a new task
* change an existing task (for example, mark completed)
* delete task
Additional requirements:
* the task list must be stored in RAM and be available at each request
* the server must respond and accept data in JSON format
* you can use the standard net/http library or try popular web libraries/frameworks (echo, chi, gorilla, etc.)*/

package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{}
var nextID = 1

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func addTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")
	var updatedTask Task
	if err := c.Bind(&updatedTask); err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
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