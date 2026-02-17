package routes

import (
	"net/http"

	"github.com/damongolding/immich-kiosk/internal/config"
	"github.com/damongolding/immich-kiosk/internal/templates/views"

	"github.com/labstack/echo/v5"
)

func Wall(baseConfig *config.Config) echo.HandlerFunc {
	return func(c *echo.Context) error {
		_ = baseConfig
		return Render(c, http.StatusOK, views.Wall())
	}
}
