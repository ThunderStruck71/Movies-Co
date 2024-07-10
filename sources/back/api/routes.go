package api

import (
	"back/internal/app/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h *handler.Handler) {
	movies := e.Group("/api/lists")
	{
		movies.GET("/", h.GetAllMovies)
		movies.GET("/:id", h.GetMovieById)
	}
}
