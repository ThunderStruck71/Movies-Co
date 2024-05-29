package handler

import (
	"back/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", Home)

	return e
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
