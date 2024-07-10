package service

import (
	"back/internal/domain/models"
	"back/internal/domain/repository"
)

type MovieService struct {
	repos repository.Movie
}

func NewMovieService(repos repository.Movie) *MovieService {
	return &MovieService{repos: repos}
}

func (ms *MovieService) GetAll() ([]models.Movie, error) {
	return ms.repos.GetAll()
}

func (ms *MovieService) GetById(movieId int) (models.Movie, error) {
	return ms.repos.GetById(movieId)
}
