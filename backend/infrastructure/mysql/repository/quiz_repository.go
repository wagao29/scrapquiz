package repository

import (
	"context"
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
