package repository

import (
	"context"
	"log"

	answerDomain "scrapquiz/domain/answer"
	utilsError "scrapquiz/utils/error"

	"scrapquiz/infrastructure/mysql/db"
	"scrapquiz/infrastructure/mysql/db/dbgen"

	"github.com/go-sql-driver/mysql"
)

type answerRepository struct{}

func NewAnswerRepository() answerDomain.AnswerRepository {
	return &answerRepository{}
}

func (r *answerRepository) Save(ctx context.Context, u *answerDomain.Answer) error {
	query := db.GetQuery(ctx)
	if err := query.InsertAnswer(ctx, dbgen.InsertAnswerParams{
		QuizID:    u.QuizID(),
		UserID:    u.UserID(),
		AnswerNum: int8(u.AnswerNum()),
	}); err != nil {
		log.Printf("[Error] AnswerRepository Save(): %v", err)

		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == utilsError.ER_DUP_ENTRY {
				return utilsError.NewBadRequestError("answer id already exists")
			}
		}
		return err
	}
	return nil
}
