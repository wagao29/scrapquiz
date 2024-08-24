package user

import (
	"net/http"

	userUseCase "scrapquiz/usecase/user"
	utilsError "scrapquiz/utils/error"
	"scrapquiz/utils/validator"

	"github.com/labstack/echo/v4"
)

type handler struct {
	uc *userUseCase.UserUseCase
}

func NewHandler(uc *userUseCase.UserUseCase) handler {
	return handler{
		uc: uc,
	}
}

func (h handler) GetUserByID(c echo.Context) error {
	inputDto := userUseCase.FindUserUseCaseInputDto{
		ID: c.Param("id"),
	}
	outputDto, err := h.uc.FindUserUseCase.Run(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}
	res := GetUserResponse{
		ID:        outputDto.ID,
		Name:      outputDto.Name,
		AvatarURL: outputDto.AvatarURL,
	}

	return c.JSON(http.StatusOK, res)
}

func (h handler) GetUsers(c echo.Context) error {
	outputDtos, err := h.uc.FindUsersUseCase.Run(c.Request().Context())
	if err != nil {
		return err
	}

	var userModels []userResponseModel
	for _, user := range outputDtos {
		userModels = append(userModels, userResponseModel{
			ID:        user.ID,
			Name:      user.Name,
			AvatarURL: user.AvatarURL,
		})
	}
	res := GetUsersResponse{userModels}

	return c.JSON(http.StatusOK, res)
}

func (h handler) PostUsers(c echo.Context) error {
	var params PostUserRequest
	err := c.Bind(&params)
	if err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	inputDto := userUseCase.SaveUserUseCaseInputDto{
		ID:        params.ID,
		Name:      params.Name,
		AvatarURL: params.AvatarURL,
	}
	outputDto, err := h.uc.SaveUserUseCase.Run(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}

	res := PostUserResponse{
		ID:        outputDto.ID,
		Name:      outputDto.Name,
		AvatarURL: outputDto.AvatarURL,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h handler) PutUser(c echo.Context) error {
	var params PutUserRequest
	err := c.Bind(&params)
	if err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		return utilsError.NewBadRequestError(err.Error())
	}

	inputDto := userUseCase.UpdateUserUseCaseInputDto{
		ID:        c.Param("id"),
		Name:      params.Name,
		AvatarURL: params.AvatarURL,
	}
	err = h.uc.UpdateUserUseCase.Run(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h handler) DeleteUserByID(c echo.Context) error {
	inputDto := userUseCase.DeleteUserUseCaseInputDto{
		ID: c.Param("id"),
	}
	err := h.uc.DeleteUserUseCase.Run(c.Request().Context(), inputDto)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
