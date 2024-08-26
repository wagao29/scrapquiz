package server

import (
	"scrapquiz/infrastructure/mysql/query_service"
	"scrapquiz/infrastructure/mysql/repository"
	"scrapquiz/interface/health_check"

	answerUseCase "scrapquiz/application/answer"
	quizUseCase "scrapquiz/application/quiz"
	userUseCase "scrapquiz/application/user"
	answerInterface "scrapquiz/interface/answer"
	quizInterface "scrapquiz/interface/quiz"
	userInterface "scrapquiz/interface/user"

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
	quizzesRoute(v1)
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

func quizzesRoute(r *echo.Group) {
	quizRepository := repository.NewQuizRepository()
	quizQueryService := query_service.NewQuizQueryService()
	answerRepository := repository.NewAnswerRepository()
	answerQueryService := query_service.NewAnswerQueryService()
	qh := quizInterface.NewHandler(quizUseCase.NewQuizUseCase(quizRepository, quizQueryService))
	ah := answerInterface.NewHandler(answerUseCase.NewAnswerUseCase(answerRepository, answerQueryService))

	group := r.Group("/quizzes")
	group.GET("/:id", qh.GetQuizByID)
	group.GET("", qh.GetQuizzes)
	group.POST("", qh.PostQuiz)
	group.DELETE("/:id", qh.DeleteQuizByID)
	group.GET("/:id/answer_counts", ah.GetAnswerCountsByAnswerID)
	group.POST("/:id/answers", ah.PostAnswer)
}
