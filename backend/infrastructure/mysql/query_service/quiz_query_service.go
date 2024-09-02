package query_service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"scrapquiz/infrastructure/mysql/db"
	"scrapquiz/infrastructure/mysql/db/dbgen"

	quizUseCase "scrapquiz/application/quiz"
	utilsError "scrapquiz/utils/error"
	utilsSQL "scrapquiz/utils/sql"
)

type quizQueryService struct{}

func NewQuizQueryService() quizUseCase.QuizQueryService {
	return &quizQueryService{}
}

func (q *quizQueryService) FetchQuizByID(
	ctx context.Context,
	id string,
) (*quizUseCase.QuizQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	dbq, err := query.FetchQuizByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utilsError.NewNotFoundError("quiz not found")
		}
		return nil, err
	}

	return &quizUseCase.QuizQueryServiceDto{
		ID:      dbq.ID,
		Content: dbq.Content,
		Options: []string{
			dbq.Option1,
			dbq.Option2,
			utilsSQL.NullStringToString(dbq.Option3),
			utilsSQL.NullStringToString(dbq.Option4),
		},
		CorrectNum:    int(dbq.CorrectNum),
		Explanation:   utilsSQL.NullStringToString(dbq.Explanation),
		UserID:        dbq.UserID,
		UserName:      dbq.UserName,
		UserAvatarURL: dbq.UserAvatarUrl,
	}, nil
}

func (q *quizQueryService) FetchQuizzesByUserID(
	ctx context.Context,
	userID string,
	limit int,
	offset int,
) ([]*quizUseCase.QuizQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	dbQuizzes, err := query.FetchQuizzesByUserID(ctx, dbgen.FetchQuizzesByUserIDParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Printf("[Error] QuizQueryService FetchQuizzesByUserID(): %v", err)
		return nil, err
	}
	if len(dbQuizzes) == 0 {
		return nil, utilsError.NewNotFoundError("quiz not found")
	}

	var dtos []*quizUseCase.QuizQueryServiceDto
	for _, dbq := range dbQuizzes {
		dtos = append(dtos, &quizUseCase.QuizQueryServiceDto{
			ID:      dbq.ID,
			Content: dbq.Content,
			Options: []string{
				dbq.Option1,
				dbq.Option2,
				utilsSQL.NullStringToString(dbq.Option3),
				utilsSQL.NullStringToString(dbq.Option4),
			},
			CorrectNum:    int(dbq.CorrectNum),
			Explanation:   utilsSQL.NullStringToString(dbq.Explanation),
			UserID:        dbq.UserID,
			UserName:      dbq.UserName,
			UserAvatarURL: dbq.UserAvatarUrl,
		})
	}
	return dtos, nil
}

func (q *quizQueryService) FetchLatestQuizzes(
	ctx context.Context,
	limit int,
	offset int,
) ([]*quizUseCase.QuizQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	dbQuizzes, err := query.FetchLatestQuizzes(ctx, dbgen.FetchLatestQuizzesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Printf("[Error] QuizQueryService FetchLatestQuizzes(): %v", err)
		return nil, err
	}
	if len(dbQuizzes) == 0 {
		return nil, utilsError.NewNotFoundError("quiz not found")
	}

	var dtos []*quizUseCase.QuizQueryServiceDto
	for _, dbq := range dbQuizzes {
		dtos = append(dtos, &quizUseCase.QuizQueryServiceDto{
			ID:      dbq.ID,
			Content: dbq.Content,
			Options: []string{
				dbq.Option1,
				dbq.Option2,
				utilsSQL.NullStringToString(dbq.Option3),
				utilsSQL.NullStringToString(dbq.Option4),
			},
			CorrectNum:    int(dbq.CorrectNum),
			Explanation:   utilsSQL.NullStringToString(dbq.Explanation),
			UserID:        dbq.UserID,
			UserName:      dbq.UserName,
			UserAvatarURL: dbq.UserAvatarUrl,
		})
	}
	return dtos, nil
}

func (q *quizQueryService) FetchQuizCounts(
	ctx context.Context,
) (int, error) {
	query := db.GetQuery(ctx)
	quizCounts, err := query.FetchQuizCounts(ctx)
	if err != nil {
		log.Printf("[Error] QuizQueryService FetchQuizCounts(): %v", err)
		return 0, err
	}
	return int(quizCounts), nil
}

func (q *quizQueryService) FetchQuizCountsByUserID(
	ctx context.Context,
	userID string,
) (int, error) {
	query := db.GetQuery(ctx)
	quizCounts, err := query.FetchQuizCountsByUserID(ctx, userID)
	if err != nil {
		log.Printf("[Error] QuizQueryService FetchQuizCountsByUserID(): %v", err)
		return 0, err
	}
	return int(quizCounts), nil
}
