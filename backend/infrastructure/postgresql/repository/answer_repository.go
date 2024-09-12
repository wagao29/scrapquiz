package repository

import (
	"context"
	"log"

	answerDomain "scrapquiz/domain/answer"

	"scrapquiz/infrastructure/postgresql/db"
	"scrapquiz/infrastructure/postgresql/db/dbgen"

	// "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
		AnswerNum: int16(u.AnswerNum()),
	}); err != nil {
		log.Printf("[Error] AnswerRepository Save(): %v", err)
		return err
	}
	return nil
}
