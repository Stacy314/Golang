package service

import (
	"time"

	"touragency/models"
	"touragency/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddOrder(tourID int, email string) (models.Order, error) {
	return s.repo.AddOrder(tourID, email)
}

func (s *Service) GetOrders() []models.Order {
	orders := s.repo.GetOrders()
	now := time.Now()
	
	for i := range orders {
		if orders[i].Date.After(now) {
			orders[i].Status = "Майбутній"
		} else if orders[i].Date.Add(24*time.Hour).After(now) {
			orders[i].Status = "Триває"
		} else {
			orders[i].Status = "Завершився"
		}
	}
	return orders
}
