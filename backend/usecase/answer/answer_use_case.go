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
}

func NewAnswerUseCase(answerRepo answerDomain.AnswerRepository) *AnswerUseCase {
	return &AnswerUseCase{
		answerRepo: answerRepo,
	}
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
