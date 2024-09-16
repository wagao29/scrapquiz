package server

import (
	"net/http"
	"os"

	utilsError "scrapquiz/utils/error"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(err error, c echo.Context) {
	switch err := err.(type) {
	case *utilsError.NotFoundError:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	case *utilsError.BadRequestError:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	default:
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.JSON(http.StatusNotFound, ErrorResponse{
					Code:    http.StatusNotFound,
					Message: "no route matches",
				})
			}
		} else {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}
}

func AuthAPIKey() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == os.Getenv("API_KEY"), nil
		},
	})
}
