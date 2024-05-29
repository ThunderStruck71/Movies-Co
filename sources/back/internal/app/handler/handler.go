package handler

import (
	"back/internal/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
