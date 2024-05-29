package service

import (
	"back/internal/domain/repository"
)

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
