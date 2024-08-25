package answer

import (
	"context"

	answerDomain "scrapquiz/domain/answer"
)

type AnswerUseCaseInputDto struct {
	UserID    string
	QuizID    string
	AnswerNum int
}
type AnswerUseCaseOutputDto struct {
	UserID    string
	QuizID    string
	AnswerNum int
}

type AnswerUseCase struct {
	answerRepo answerDomain.AnswerRepository
	answerQS   AnswerQueryService
}

func NewAnswerUseCase(
	answerRepo answerDomain.AnswerRepository,
	answerQS AnswerQueryService,
) *AnswerUseCase {
	return &AnswerUseCase{
		answerRepo: answerRepo,
		answerQS:   answerQS,
	}
}

func (uc *AnswerUseCase) FetchAnswerCountsByQuizID(ctx context.Context, id string) (*AnswerQueryServiceDto, error) {
	quiz, err := uc.answerQS.FetchAnswerCountsByQuizID(ctx, id)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}

func (uc *AnswerUseCase) Save(ctx context.Context, dto AnswerUseCaseInputDto) (*AnswerUseCaseOutputDto, error) {
	answer, err := answerDomain.NewAnswer(dto.QuizID, dto.UserID, dto.AnswerNum)
	if err != nil {
		return nil, err
	}
	if err = uc.answerRepo.Save(ctx, answer); err != nil {
		return nil, err
	}
	return &AnswerUseCaseOutputDto{
		QuizID:    answer.QuizID(),
		UserID:    answer.UserID(),
		AnswerNum: answer.AnswerNum(),
	}, nil
}
