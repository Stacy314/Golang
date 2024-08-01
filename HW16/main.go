/*Add to the REST API implementation in the task HW14 (Layered REST) ​​or HW10 (REST) ​​data storage in a relational database.
Additional conditions:
* use one of the popular SQL RDBMS (MySQL, Postgres,...)
* add a docker-compose file to run the database locally in Docker
* use any library to connect to a SQL database
* describe the database scheme (ideally, in the form of migration)*/

package main

import (
    "context"
    "log"

    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "HW16/handlers"
    "HW16/service"
)

func main() {
    db, err := pgxpool.Connect(context.Background(), "postgres://user:password@localhost:5432/tasksdb")
    if err != nil {
        log.Fatal("Unable to connect to database:", err)
    }
    defer db.Close()

    taskService := service.NewTaskService(db)

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Set("taskService", taskService)
            return next(c)
        }
    })

    e.GET("/tasks", handlers.GetTasks)
    e.POST("/tasks", handlers.AddTask)
    e.PUT("/tasks/:id", handlers.UpdateTask)
    e.DELETE("/tasks/:id", handlers.DeleteTask)

    e.Logger.Fatal(e.Start(":8080"))
}
