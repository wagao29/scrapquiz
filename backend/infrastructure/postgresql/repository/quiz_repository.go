package repository

import (
	"context"
	"log"

	quizDomain "scrapquiz/domain/quiz"
	utilsSQL "scrapquiz/utils/sql"

	"scrapquiz/infrastructure/postgresql/db"
	"scrapquiz/infrastructure/postgresql/db/dbgen"
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
		CorrectNum:  int16(q.CorrectNum()),
		Explanation: utilsSQL.StringToNullString(q.Explanation()),
	}); err != nil {
		log.Printf("[Error] QuizRepository Save(): %v", err)
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
