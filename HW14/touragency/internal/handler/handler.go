package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"touragency/internal/service"
	"touragency/models"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetTours(c echo.Context) error {
	tours := h.svc.GetTours()
	return c.JSON(http.StatusOK, tours)
}

func (h *Handler) AddOrder(c echo.Context) error {
	var input struct {
		TourID int    `json:"tour_id"`
		Email  string `json:"email"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Невірні дані"})
	}

	order, err := h.svc.AddOrder(input.TourID, input.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *Handler) GetOrders(c echo.Context) error {
	orders := h.svc.GetOrders()
	return c.JSON(http.StatusOK, orders)
}