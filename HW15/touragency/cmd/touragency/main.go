/*Create a REST API for three endpoints on a free topic based on a three-layer architecture. 
It shouldn't be just CRUD, you need to add some business logic. User authorization may not be 
implemented.
Example: API for a travel agency website
* receiving a list of available tours (name, description, price, mode of transport)
* order a tour. In addition to saving the order, we simulate sending an email (log that it 
has been sent)
* list of ordered tours: you can add additional properties/business logic for ordered tours. 
For example, the status of the tour depending on the current date (future, soon, ongoing, completed)*/

package main

import (
	"touragency/internal/handler"
	"touragency/internal/repository"
	"touragency/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tours", h.GetTours)
	e.POST("/orders", h.AddOrder)
	e.GET("/orders", h.GetOrders)

	e.Logger.Fatal(e.Start(":8080"))
}