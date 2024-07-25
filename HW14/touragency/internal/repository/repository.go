package repository

import (
	"errors"
	"sync"

	"touragency/models"
)

type Repository struct {
	tours  []models.Tour
	orders []models.Order
	mu     sync.Mutex
	nextID int
}

func NewRepository() *Repository {
	return &Repository{
		tours: []models.Tour{
			{ID: 1, Title: "Париж", Description: "Тур до Парижу", Price: 300.0, Transport: "Літак"},
			{ID: 2, Title: "Лондон", Description: "Тур до Лондону", Price: 400.0, Transport: "Літак"},
		},
		orders: []models.Order{},
		nextID: 1,
	}
}

func (r *Repository) GetTours() []models.Tour {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	return r.tours
}

func (r *Repository) AddOrder(tourID int, email string) (models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, tour := range r.tours {
		if tour.ID == tourID {
			order := models.Order{
				ID:     r.nextID,
				TourID: tourID,
				Email:  email,
				Date:   time.Now(),
			}
			r.orders = append(r.orders, order)
			r.nextID++
			return order, nil
		}
	}
	return models.Order{}, errors.New("Тур не знайдено")
}

func (r *Repository) GetOrders() []models.Order {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	return r.orders
}
