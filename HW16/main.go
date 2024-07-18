package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *pgxpool.Pool

func getTasks(c echo.Context) error {
	rows, err := db.Query(context.Background(), "SELECT id, title, completed FROM tasks")
	if err != nil {
		return err
	}
	defer rows.Close()

	tasks := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return err
		}
		tasks = append(tasks, task)
	}

	return c.JSON(http.StatusOK, tasks)
}

func addTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return err
	}
	err := db.QueryRow(context.Background(), 
		"INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id", 
		task.Title, task.Completed).Scan(&task.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTask Task
	if err := c.Bind(&updatedTask); err != nil {
		return err
	}

	commandTag, err := db.Exec(context.Background(), 
		"UPDATE tasks SET title=$1, completed=$2 WHERE id=$3", 
		updatedTask.Title, updatedTask.Completed, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
	}

	updatedTask.ID = id
	return c.JSON(http.StatusOK, updatedTask)
}

func deleteTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	commandTag, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	var err error
	db, err = pgxpool.Connect(context.Background(), "postgres://user:password@localhost:5432/tasksdb")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", addTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}