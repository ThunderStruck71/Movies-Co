package repository

import (
	"back/internal/domain/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MoviePostgres struct {
	db *sqlx.DB
}

func NewMoviePostgres(db *sqlx.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (mp *MoviePostgres) GetAll() ([]models.Movie, error) {
	var movies []models.Movie
	query := fmt.Sprintf(`select mt.original_title, mt.overview, mt.poster_path, mt.release_date, mt.is_watched from %s mt`, moviesTable)

	if err := mp.db.Select(&movies, query); err != nil {
		return nil, err
	}

	return movies, nil
}

func (mp *MoviePostgres) GetById(movieId int) (models.Movie, error) {
	var movie models.Movie
	query := fmt.Sprintf(`select mt.original_title, mt.overview, mt.poster_path, mt.release_date, mt.is_watched from %s mt where mt.id = $1`, moviesTable)

	err := mp.db.Get(&movie, query, movieId)

	return movie, err
}
