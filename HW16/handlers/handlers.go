package handlers

import (
    "context"
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "HW16/service"
)

func GetTasks(c echo.Context) error {
    taskService := c.Get("taskService").(*service.TaskService)
    tasks, err := taskService.GetTasks(context.Background())
    if err != nil {
        return err
    }
    return c.JSON(http.StatusOK, tasks)
}

func AddTask(c echo.Context) error {
    taskService := c.Get("taskService").(*service.TaskService)
    var task service.Task
    if err := c.Bind(&task); err != nil {
        return err
    }
    if err := taskService.AddTask(context.Background(), &task); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, task)
}

func UpdateTask(c echo.Context) error {
    taskService := c.Get("taskService").(*service.TaskService)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task ID"})
    }

    var updatedTask service.Task
    if err := c.Bind(&updatedTask); err != nil {
        return err
    }
    updatedTask.ID = id

    if err := taskService.UpdateTask(context.Background(), &updatedTask); err != nil {
        if err == service.ErrNoRows {
            return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
        }
        return err
    }

    return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
    taskService := c.Get("taskService").(*service.TaskService)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid task ID"})
    }

    if err := taskService.DeleteTask(context.Background(), id); err != nil {
        if err == service.ErrNoRows {
            return c.JSON(http.StatusNotFound, echo.Map{"error": "Task not found"})
        }
        return err
    }
    return c.NoContent(http.StatusNoContent)
}