package models

type Movie struct {
	Id            int    `json:"id" db:"id" binding:"required"`
	OriginalTitle string `json:"original_title" db:"original_title" binding:"required"`
	Overview      string `json:"overview" db:"overview"`
	PosterPath    string `json:"poster_path" db:"poster_path"`
	ReleaseDate   string `json:"release_date" db:"release_date"`
	IsWatched     bool   `json:"is_watched" db:"is_watched"`
}
