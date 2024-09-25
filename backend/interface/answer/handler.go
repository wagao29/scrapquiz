package answer

import (
	"net/http"

	answerUseCase "scrapquiz/application/answer"
	utilsError "scrapquiz/utils/error"
	"scrapquiz/utils/validator"

	"github.com/labstack/echo/v4"
)

type handler struct {
	uc *answerUseCase.AnswerUseCase
}

func NewHandler(uc *answerUseCase.AnswerUseCase) handler {
	return handler{
		uc: uc,
	}
}

func (h handler) GetAnswerCountsByAnswerID(c echo.Context) error {
	outputDto, err := h.uc.FetchAnswerCountsByQuizID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	res := FetchAnswerCountsResponse{
		AnswerCounts: []int{
			outputDto.AnswerNum1Count,
			outputDto.AnswerNum2Count,
			outputDto.AnswerNum3Count,
			outputDto.AnswerNum4Count,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) PostAnswer(c echo.Context) error {
	var params PostAnswerRequest
	err := c.Bind(&params)
	if err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	inputDto := answerUseCase.AnswerUseCaseInputDto{
		QuizID:    c.Param("id"),
		UserID:    params.UserID,
		AnswerNum: params.AnswerNum,
	}
	outputDto, err := h.uc.Save(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}

	res := PostAnswerResponse{
		QuizID:    outputDto.QuizID,
		UserID:    outputDto.UserID,
		AnswerNum: outputDto.AnswerNum,
	}

	return c.JSON(http.StatusCreated, res)
}
