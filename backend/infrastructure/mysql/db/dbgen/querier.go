// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbgen

import (
	"context"
)

type Querier interface {
	DeleteQuiz(ctx context.Context, id string) error
	DeleteUser(ctx context.Context, id string) error
	FetchAnswerCountsByQuizID(ctx context.Context, quizID string) ([]FetchAnswerCountsByQuizIDRow, error)
	FetchLatestQuizzes(ctx context.Context, arg FetchLatestQuizzesParams) ([]FetchLatestQuizzesRow, error)
	FetchQuizByID(ctx context.Context, id string) (FetchQuizByIDRow, error)
	FetchQuizCounts(ctx context.Context) (int64, error)
	FetchQuizCountsByUserID(ctx context.Context, userID string) (int64, error)
	FetchQuizzesByUserID(ctx context.Context, arg FetchQuizzesByUserIDParams) ([]FetchQuizzesByUserIDRow, error)
	InsertAnswer(ctx context.Context, arg InsertAnswerParams) error
	InsertQuiz(ctx context.Context, arg InsertQuizParams) error
	InsertUser(ctx context.Context, arg InsertUserParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UserFindAll(ctx context.Context) ([]User, error)
	UserFindByID(ctx context.Context, id string) (User, error)
}

var _ Querier = (*Queries)(nil)
