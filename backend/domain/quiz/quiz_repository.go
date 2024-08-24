package quiz

import (
	"context"
)

type QuizRepository interface {
	Save(ctx context.Context, quiz *Quiz) error
	Delete(ctx context.Context, id string) error
}
