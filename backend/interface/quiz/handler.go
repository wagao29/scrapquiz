package quiz

import (
	"net/http"
	"strconv"

	quizUseCase "scrapquiz/application/quiz"
	utilsError "scrapquiz/utils/error"
	"scrapquiz/utils/validator"

	"github.com/labstack/echo/v4"
)

type handler struct {
	uc *quizUseCase.QuizUseCase
}

func NewHandler(uc *quizUseCase.QuizUseCase) handler {
	return handler{
		uc: uc,
	}
}

func (h handler) GetQuizByID(c echo.Context) error {
	outputDto, err := h.uc.FetchQuizByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	res := FetchQuizResponse{
		ID:            outputDto.ID,
		Content:       outputDto.Content,
		Options:       outputDto.Options,
		CorrectNum:    outputDto.CorrectNum,
		Explanation:   outputDto.Explanation,
		UserID:        outputDto.UserID,
		UserName:      outputDto.UserName,
		UserAvatarURL: outputDto.UserAvatarURL,
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetQuizzes(c echo.Context) error {
	userID := c.QueryParam("user_id")

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return utilsError.NewBadRequestError("limit param is invalid")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return utilsError.NewBadRequestError("offset param is invalid")
	}

	var outputDtos []*quizUseCase.QuizQueryServiceDto
	if userID != "" {
		outputDtos, err = h.uc.FetchQuizzesByUserID(c.Request().Context(), userID, limit, offset)
		if err != nil {
			return err
		}
	} else {
		outputDtos, err = h.uc.FetchLatestQuizzes(c.Request().Context(), limit, offset)
		if err != nil {
			return err
		}
	}

	var fetchQuizzes []FetchQuizResponse
	for _, outputDto := range outputDtos {
		fetchQuizzes = append(fetchQuizzes, FetchQuizResponse{
			ID:            outputDto.ID,
			Content:       outputDto.Content,
			Options:       outputDto.Options,
			CorrectNum:    outputDto.CorrectNum,
			Explanation:   outputDto.Explanation,
			UserID:        outputDto.UserID,
			UserName:      outputDto.UserName,
			UserAvatarURL: outputDto.UserAvatarURL,
		})
	}
	res := FetchQuizzesResponse{fetchQuizzes}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetQuizCounts(c echo.Context) error {
	userID := c.QueryParam("user_id")

	var (
		quizCounts int
		err        error
	)
	if len(userID) == 0 {
		quizCounts, err = h.uc.FetchQuizCounts(c.Request().Context())

	} else {
		quizCounts, err = h.uc.FetchQuizCountsByUserID(c.Request().Context(), userID)
	}
	if err != nil {
		return err
	}
	res := FetchQuizCountsResponse{quizCounts}
	return c.JSON(http.StatusOK, res)
}

func (h handler) PostQuiz(c echo.Context) error {
	var params PostQuizRequest
	err := c.Bind(&params)
	if err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	inputDto := quizUseCase.QuizUseCaseSaveInputDto{
		UserID:      params.UserID,
		Content:     params.Content,
		Options:     params.Options,
		CorrectNum:  params.CorrectNum,
		Explanation: params.Explanation,
	}
	outputDto, err := h.uc.Save(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}

	res := PostQuizResponse{
		ID:          outputDto.ID,
		UserID:      outputDto.UserID,
		Content:     outputDto.Content,
		Options:     outputDto.Options,
		CorrectNum:  outputDto.CorrectNum,
		Explanation: outputDto.Explanation,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h handler) DeleteQuizByID(c echo.Context) error {
	err := h.uc.Delete(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
