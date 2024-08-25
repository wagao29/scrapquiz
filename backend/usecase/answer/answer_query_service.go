package answer

import "context"

type AnswerQueryServiceDto struct {
	AnswerNum1Count int
	AnswerNum2Count int
	AnswerNum3Count int
	AnswerNum4Count int
}

type AnswerQueryService interface {
	FetchAnswerCountsByQuizID(ctx context.Context, quizID string) (*AnswerQueryServiceDto, error)
}
