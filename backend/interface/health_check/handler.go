package health_check

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	res := HealthCheckResponse{
		Status: "ok",
	}
	return c.JSON(http.StatusOK, res)
}
