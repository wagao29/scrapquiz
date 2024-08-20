package server

import (
	"scrapquiz/infrastructure/mysql/repository"
	"scrapquiz/interface/health_check"
	userInterface "scrapquiz/interface/user"
	userUseCase "scrapquiz/usecase/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler

	v1 := e.Group("/v1")
	v1.GET("/health", health_check.HealthCheck)

	userRoute(v1)
}

func userRoute(r *echo.Group) {
	userRepository := repository.NewUserRepository()
	h := userInterface.NewHandler(userUseCase.NewUserUseCase(userRepository))
	group := r.Group("/users")
	group.GET("/:id", h.GetUserByID)
	group.GET("", h.GetUsers)
	group.POST("", h.PostUsers)
	group.PUT("/:id", h.PutUser)
	group.DELETE("/:id", h.DeleteUserByID)
}
