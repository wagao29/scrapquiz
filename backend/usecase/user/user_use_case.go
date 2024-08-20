package user

import (
	"context"

	userDomain "scrapquiz/domain/user"
)

// UseCase types
type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}
type FindUsersUseCase struct {
	userRepo userDomain.UserRepository
}
type SaveUserUseCase struct {
	userRepo userDomain.UserRepository
}
type UpdateUserUseCase struct {
	userRepo userDomain.UserRepository
}
type DeleteUserUseCase struct {
	userRepo userDomain.UserRepository
}
type UserUseCase struct {
	FindUserUseCase   *FindUserUseCase
	FindUsersUseCase  *FindUsersUseCase
	SaveUserUseCase   *SaveUserUseCase
	UpdateUserUseCase *UpdateUserUseCase
	DeleteUserUseCase *DeleteUserUseCase
}

// DTO types
type userUseCaseDto struct {
	ID        string
	Name      string
	AvatarURL string
}
type FindUserUseCaseInputDto struct {
	ID string
}
type DeleteUserUseCaseInputDto struct {
	ID string
}
type (
	FindUserUseCaseOutputDto  = userUseCaseDto
	SaveUserUseCaseInputDto   = userUseCaseDto
	SaveUserUseCaseOutputDto  = userUseCaseDto
	UpdateUserUseCaseInputDto = userUseCaseDto
	FindUsersUseCaseOutputDto = []*userUseCaseDto
)

// Factory functions
func NewUserUseCase(userRepo userDomain.UserRepository) *UserUseCase {
	return &UserUseCase{
		newFindUserUseCase(userRepo),
		newFindUsersUseCase(userRepo),
		newSaveUserUseCase(userRepo),
		newUpdateUserUseCase(userRepo),
		newDeleteUserUseCase(userRepo),
	}
}

func newFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

func newFindUsersUseCase(
	userRepo userDomain.UserRepository,
) *FindUsersUseCase {
	return &FindUsersUseCase{
		userRepo: userRepo,
	}
}

func newSaveUserUseCase(
	userRepo userDomain.UserRepository,
) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepo: userRepo,
	}
}

func newUpdateUserUseCase(
	userRepo userDomain.UserRepository,
) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepo: userRepo,
	}
}

func newDeleteUserUseCase(
	userRepo userDomain.UserRepository,
) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepo: userRepo,
	}
}

// Run Methods
func (uc *FindUserUseCase) Run(ctx context.Context, dto FindUserUseCaseInputDto) (*FindUserUseCaseOutputDto, error) {
	user, err := uc.userRepo.FindByID(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseOutputDto{
		ID:        user.ID(),
		Name:      user.Name(),
		AvatarURL: user.AvatarURL(),
	}, nil
}

func (uc *FindUsersUseCase) Run(ctx context.Context) (FindUsersUseCaseOutputDto, error) {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var dtoUsers []*userUseCaseDto
	for _, user := range users {
		dtoUsers = append(dtoUsers, &userUseCaseDto{
			ID:        user.ID(),
			Name:      user.Name(),
			AvatarURL: user.AvatarURL(),
		})
	}

	return dtoUsers, nil
}

func (uc *SaveUserUseCase) Run(ctx context.Context, dto SaveUserUseCaseInputDto) (*SaveUserUseCaseOutputDto, error) {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.AvatarURL)
	if err != nil {
		return nil, err
	}
	if err = uc.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}
	return &SaveUserUseCaseOutputDto{
		ID:        user.ID(),
		Name:      user.Name(),
		AvatarURL: user.AvatarURL(),
	}, nil
}

func (uc *UpdateUserUseCase) Run(ctx context.Context, dto SaveUserUseCaseInputDto) error {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.AvatarURL)
	if err != nil {
		return err
	}
	if err = uc.userRepo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (uc *DeleteUserUseCase) Run(ctx context.Context, dto DeleteUserUseCaseInputDto) error {
	if err := uc.userRepo.Delete(ctx, dto.ID); err != nil {
		return err
	}
	return nil
}
