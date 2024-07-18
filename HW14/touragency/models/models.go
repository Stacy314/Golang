package models

import "time"

type Tour struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Transport   string  `json:"transport"`
}

type Order struct {
	ID      int       `json:"id"`
	TourID  int       `json:"tour_id"`
	Email   string    `json:"email"`
	Date    time.Time `json:"date"`
	Status  string    `json:"status"`
}