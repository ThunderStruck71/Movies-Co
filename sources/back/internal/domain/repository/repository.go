package repository

import (
	"back/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type Movie interface {
	GetAll() ([]models.Movie, error)
	GetById(movieId int) (models.Movie, error)
}

type Repository struct {
	Movie
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Movie: NewMoviePostgres(db),
	}
}
