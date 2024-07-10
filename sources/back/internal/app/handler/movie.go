package handler

import (
	"back/internal/domain/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type customError struct {
	Message string `json:"message"`
}

type getAllMoviesResponse struct {
	Data []models.Movie `json:"data"`
}

func newErrorResponse(ctx echo.Context, statusCode int, message string) {
	log.Fatalf(message)
	ctx.JSON(statusCode, customError{Message: message})
}

func (h *Handler) GetAllMovies(ctx echo.Context) error {
	data, err := h.services.Movie.GetAll()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, getAllMoviesResponse{
		Data: data,
	})
}

func (h *Handler) GetMovieById(ctx echo.Context) error {
	movieId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return err
	}

	data, err := h.services.Movie.GetById(movieId)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
