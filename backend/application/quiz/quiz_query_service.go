package quiz

import (
	"context"
	"time"
)

type QuizQueryServiceDto struct {
	ID            string
	Content       string
	Options       []string
	CorrectNum    int
	Explanation   string
	UserID        string
	UserName      string
	UserAvatarURL string
	CreatedAt     time.Time
}

type QuizQueryService interface {
	FetchQuizByID(ctx context.Context, userID string) (*QuizQueryServiceDto, error)
	FetchQuizzesByUserID(ctx context.Context, userID string, limit int, offset int) ([]*QuizQueryServiceDto, error)
	FetchLatestQuizzes(ctx context.Context, limit int, offset int) ([]*QuizQueryServiceDto, error)
	FetchRandomQuizzes(ctx context.Context, limit int) ([]*QuizQueryServiceDto, error)
	FetchQuizCounts(ctx context.Context) (int, error)
	FetchQuizCountsByUserID(ctx context.Context, userID string) (int, error)
}
