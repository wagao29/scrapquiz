package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	userDomain "scrapquiz/domain/user"
	utilsError "scrapquiz/utils/error"

	"scrapquiz/infrastructure/mysql/db"
	"scrapquiz/infrastructure/mysql/db/dbgen"

	"github.com/go-sql-driver/mysql"
)

type userRepository struct{}

func NewUserRepository() userDomain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	query := db.GetQuery(ctx)
	user, err := query.UserFindByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utilsError.NewNotFoundError("user not found")
		}
		return nil, err
	}
	domainUser, err := userDomain.NewUser(
		user.ID,
		user.Name,
		user.AvatarUrl,
	)
	if err != nil {
		log.Printf("[Error] UserRepository FindByID(): %v", err)
		return nil, err
	}
	return domainUser, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	query := db.GetQuery(ctx)
	queryUsers, err := query.UserFindAll(ctx)
	if err != nil {
		log.Printf("[Error] UserRepository FindAll(): %v", err)
		return nil, err
	}
	if len(queryUsers) == 0 {
		return nil, utilsError.NewNotFoundError("user not found")
	}

	var users []*userDomain.User
	for _, qu := range queryUsers {
		du, err := userDomain.NewUser(
			qu.ID,
			qu.Name,
			qu.AvatarUrl,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, du)
	}

	return users, nil
}

func (r *userRepository) Save(ctx context.Context, u *userDomain.User) error {
	query := db.GetQuery(ctx)
	if err := query.InsertUser(ctx, dbgen.InsertUserParams{
		ID:        u.ID(),
		Name:      u.Name(),
		AvatarUrl: u.AvatarURL(),
	}); err != nil {
		log.Printf("[Error] UserRepository Save(): %v", err)

		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == utilsError.ER_DUP_ENTRY {
				return utilsError.NewBadRequestError("user id already exists")
			}
		}
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, u *userDomain.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpdateUser(ctx, dbgen.UpdateUserParams{
		ID:        u.ID(),
		Name:      u.Name(),
		AvatarUrl: u.AvatarURL(),
	}); err != nil {
		log.Printf("[Error] UserRepository Update(): %v", err)
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := db.GetQuery(ctx)
	if err := query.DeleteUser(ctx, id); err != nil {
		log.Printf("[Error] UserRepository Delete(): %v", err)
		return err
	}
	return nil
}
