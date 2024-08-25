package answer

import (
	"context"
)

type AnswerRepository interface {
	Save(ctx context.Context, quiz *Answer) error
}
