package quiz

import (
	"context"
)

type QuizRepository interface {
	FindByID(ctx context.Context, id string) (*Quiz, error)
	FindByUserID(ctx context.Context, userID string, limit int, offset int) ([]*Quiz, error)
	FindLatest(ctx context.Context, limit int, offset int) ([]*Quiz, error)
	Save(ctx context.Context, quiz *Quiz) error
	Delete(ctx context.Context, id string) error
}
