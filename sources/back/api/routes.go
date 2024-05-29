package api

import (
	"back/internal/app/handler"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h *handler.Handler) {
	e.GET("/", h.Home)
}
