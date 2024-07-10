package service

import (
	"back/internal/domain/models"
	"back/internal/domain/repository"
)

type Movie interface {
	GetAll() ([]models.Movie, error)
	GetById(movieId int) (models.Movie, error)
}

type Service struct {
	Movie
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Movie: NewMovieService(repos.Movie),
	}
}
