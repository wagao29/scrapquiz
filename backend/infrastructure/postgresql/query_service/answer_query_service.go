package query_service

import (
	"context"
	"database/sql"
	"errors"

	"scrapquiz/infrastructure/postgresql/db"

	answerUseCase "scrapquiz/application/answer"
	utilsError "scrapquiz/utils/error"
)

type answerQueryService struct{}

func NewAnswerQueryService() answerUseCase.AnswerQueryService {
	return &answerQueryService{}
}

func (q *answerQueryService) FetchAnswerCountsByQuizID(
	ctx context.Context,
	id string,
) (*answerUseCase.AnswerQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	answers, err := query.FetchAnswerCountsByQuizID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utilsError.NewNotFoundError("answer not found")
		}
		return nil, err
	}

	var dto = &answerUseCase.AnswerQueryServiceDto{}
	for _, answer := range answers {
		count := int(answer.Count)
		switch answer.AnswerNum {
		case 1:
			dto.AnswerNum1Count = count
		case 2:
			dto.AnswerNum2Count = count
		case 3:
			dto.AnswerNum3Count = count
		case 4:
			dto.AnswerNum4Count = count
		}
	}

	return dto, nil
}
