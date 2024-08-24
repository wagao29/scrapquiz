package user

import (
	"context"

	userDomain "scrapquiz/domain/user"
)

type UserUseCaseInputDto struct {
	ID        string
	Name      string
	AvatarURL string
}
type UserUseCaseOutputDto struct {
	ID        string
	Name      string
	AvatarURL string
}

type UserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewUserUseCase(userRepo userDomain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) FindByID(ctx context.Context, id string) (*UserUseCaseOutputDto, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &UserUseCaseOutputDto{
		ID:        user.ID(),
		Name:      user.Name(),
		AvatarURL: user.AvatarURL(),
	}, nil
}

func (uc *UserUseCase) FindAll(ctx context.Context) ([]*UserUseCaseOutputDto, error) {
	users, err := uc.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var dtoUsers []*UserUseCaseOutputDto
	for _, user := range users {
		dtoUsers = append(dtoUsers, &UserUseCaseOutputDto{
			ID:        user.ID(),
			Name:      user.Name(),
			AvatarURL: user.AvatarURL(),
		})
	}

	return dtoUsers, nil
}

func (uc *UserUseCase) Save(ctx context.Context, dto UserUseCaseInputDto) (*UserUseCaseOutputDto, error) {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.AvatarURL)
	if err != nil {
		return nil, err
	}
	if err = uc.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}
	return &UserUseCaseOutputDto{
		ID:        user.ID(),
		Name:      user.Name(),
		AvatarURL: user.AvatarURL(),
	}, nil
}

func (uc *UserUseCase) Update(ctx context.Context, dto UserUseCaseInputDto) error {
	user, err := userDomain.NewUser(dto.ID, dto.Name, dto.AvatarURL)
	if err != nil {
		return err
	}
	if err = uc.userRepo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) Delete(ctx context.Context, id string) error {
	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
