package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	quizDomain "scrapquiz/domain/quiz"
	utilsError "scrapquiz/utils/error"
	utilsSQL "scrapquiz/utils/sql"

	"scrapquiz/infrastructure/mysql/db"
	"scrapquiz/infrastructure/mysql/db/dbgen"

	"github.com/go-sql-driver/mysql"
)

type quizRepository struct{}

func NewQuizRepository() quizDomain.QuizRepository {
	return &quizRepository{}
}

func (r *quizRepository) FindByID(ctx context.Context, id string) (*quizDomain.Quiz, error) {
	query := db.GetQuery(ctx)
	quiz, err := query.FindQuizByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utilsError.NewNotFoundError("quiz not found")
		}
		return nil, err
	}
	domainQuiz, err := quizDomain.Reconstruct(
		quiz.ID,
		quiz.UserID,
		quiz.Content,
		[]string{
			quiz.Option1,
			quiz.Option2,
			utilsSQL.NullStringToString(quiz.Option3),
			utilsSQL.NullStringToString(quiz.Option4),
		},
		int(quiz.CorrectNum),
		utilsSQL.NullStringToString(quiz.Explanation),
	)
	if err != nil {
		log.Printf("[Error] QuizRepository FindByID(): %v", err)
		return nil, err
	}
	return domainQuiz, nil
}

func (r *quizRepository) FindByUserID(
	ctx context.Context,
	userID string,
	limit int,
	offset int,
) ([]*quizDomain.Quiz, error) {
	query := db.GetQuery(ctx)
	dbQuizzes, err := query.FindQuizzesByUserID(
		ctx,
		dbgen.FindQuizzesByUserIDParams{
			UserID: userID,
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		log.Printf("[Error] QuizRepository FindByUserID(): %v", err)
		return nil, err
	}
	if len(dbQuizzes) == 0 {
		return nil, utilsError.NewNotFoundError("quiz not found")
	}

	return mapToDomainQuiz(dbQuizzes)
}

func (r *quizRepository) FindLatest(
	ctx context.Context,
	limit int,
	offset int,
) ([]*quizDomain.Quiz, error) {
	query := db.GetQuery(ctx)
	dbQuizzes, err := query.FindLatestQuizzes(
		ctx,
		dbgen.FindLatestQuizzesParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
	if err != nil {
		log.Printf("[Error] QuizRepository FindByUserID(): %v", err)
		return nil, err
	}
	if len(dbQuizzes) == 0 {
		return nil, utilsError.NewNotFoundError("quiz not found")
	}

	return mapToDomainQuiz(dbQuizzes)
}

func (r *quizRepository) Save(ctx context.Context, q *quizDomain.Quiz) error {
	query := db.GetQuery(ctx)
	if err := query.InsertQuiz(ctx, dbgen.InsertQuizParams{
		ID:          q.ID(),
		UserID:      q.UserID(),
		Content:     q.Content(),
		Option1:     q.Option1(),
		Option2:     q.Option2(),
		Option3:     utilsSQL.StringToNullString(q.Option3()),
		Option4:     utilsSQL.StringToNullString(q.Option4()),
		CorrectNum:  int8(q.CorrectNum()),
		Explanation: utilsSQL.StringToNullString(q.Explanation()),
	}); err != nil {
		log.Printf("[Error] QuizRepository Save(): %v", err)

		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == utilsError.ER_DUP_ENTRY {
				return utilsError.NewBadRequestError("quiz id already exists")
			}
		}
		return err
	}
	return nil
}

func (r *quizRepository) Delete(ctx context.Context, id string) error {
	query := db.GetQuery(ctx)
	if err := query.DeleteQuiz(ctx, id); err != nil {
		log.Printf("[Error] QuizRepository Delete(): %v", err)
		return err
	}
	return nil
}

func mapToDomainQuiz(dbQuizzes []dbgen.Quiz) ([]*quizDomain.Quiz, error) {
	var domainQuizzes []*quizDomain.Quiz
	for _, dbq := range dbQuizzes {
		du, err := quizDomain.Reconstruct(
			dbq.ID,
			dbq.UserID,
			dbq.Content,
			[]string{
				dbq.Option1,
				dbq.Option2,
				utilsSQL.NullStringToString(dbq.Option3),
				utilsSQL.NullStringToString(dbq.Option4),
			},
			int(dbq.CorrectNum),
			utilsSQL.NullStringToString(dbq.Explanation),
		)
		if err != nil {
			return nil, err
		}
		domainQuizzes = append(domainQuizzes, du)
	}
	return domainQuizzes, nil
}
