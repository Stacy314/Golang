package service

import (
	"errors"
	"log"
	"touragency/models"
	"touragency/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTours() []models.Tour {
	return s.repo.GetTours()
}

func (s *Service) AddOrder(tourID int, email string) (models.Order, error) {
	order, err := s.repo.AddOrder(tourID, email)
	if err != nil {
		return models.Order{}, err
	}
	log.Printf("Email відправлено на %s для замовлення %d", email, order.ID)
	return order, nil
}

func (s *Service) GetOrders() []models.Order {
	return s.repo.GetOrders()
}