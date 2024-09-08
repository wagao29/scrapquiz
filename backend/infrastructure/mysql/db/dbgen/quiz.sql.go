// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: quiz.sql

package dbgen

import (
	"context"
	"database/sql"
	"time"
)

const deleteQuiz = `-- name: DeleteQuiz :exec
DELETE FROM quizzes
WHERE
  id = ?
`

func (q *Queries) DeleteQuiz(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteQuiz, id)
	return err
}

const fetchLatestQuizzes = `-- name: FetchLatestQuizzes :many
SELECT
  quizzes.id,
  quizzes.content,
  quizzes.option_1,
  quizzes.option_2,
  quizzes.option_3,
  quizzes.option_4,
  quizzes.correct_num,
  quizzes.explanation,
  quizzes.user_id,
  users.name AS user_name,
  users.avatar_url AS user_avatar_url,
  quizzes.created_at
FROM
  quizzes
INNER JOIN users ON quizzes.user_id = users.id
ORDER BY id DESC
LIMIT ? OFFSET ?
`

type FetchLatestQuizzesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type FetchLatestQuizzesRow struct {
	ID            string         `json:"id"`
	Content       string         `json:"content"`
	Option1       string         `json:"option_1"`
	Option2       string         `json:"option_2"`
	Option3       sql.NullString `json:"option_3"`
	Option4       sql.NullString `json:"option_4"`
	CorrectNum    int8           `json:"correct_num"`
	Explanation   sql.NullString `json:"explanation"`
	UserID        string         `json:"user_id"`
	UserName      string         `json:"user_name"`
	UserAvatarUrl string         `json:"user_avatar_url"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (q *Queries) FetchLatestQuizzes(ctx context.Context, arg FetchLatestQuizzesParams) ([]FetchLatestQuizzesRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchLatestQuizzes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchLatestQuizzesRow{}
	for rows.Next() {
		var i FetchLatestQuizzesRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Option1,
			&i.Option2,
			&i.Option3,
			&i.Option4,
			&i.CorrectNum,
			&i.Explanation,
			&i.UserID,
			&i.UserName,
			&i.UserAvatarUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchQuizByID = `-- name: FetchQuizByID :one
SELECT
  quizzes.id,
  quizzes.content,
  quizzes.option_1,
  quizzes.option_2,
  quizzes.option_3,
  quizzes.option_4,
  quizzes.correct_num,
  quizzes.explanation,
  quizzes.user_id,
  users.name AS user_name,
  users.avatar_url AS user_avatar_url,
  quizzes.created_at
FROM
  quizzes
INNER JOIN users ON quizzes.user_id = users.id
WHERE
  quizzes.id = ?
`

type FetchQuizByIDRow struct {
	ID            string         `json:"id"`
	Content       string         `json:"content"`
	Option1       string         `json:"option_1"`
	Option2       string         `json:"option_2"`
	Option3       sql.NullString `json:"option_3"`
	Option4       sql.NullString `json:"option_4"`
	CorrectNum    int8           `json:"correct_num"`
	Explanation   sql.NullString `json:"explanation"`
	UserID        string         `json:"user_id"`
	UserName      string         `json:"user_name"`
	UserAvatarUrl string         `json:"user_avatar_url"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (q *Queries) FetchQuizByID(ctx context.Context, id string) (FetchQuizByIDRow, error) {
	row := q.db.QueryRowContext(ctx, fetchQuizByID, id)
	var i FetchQuizByIDRow
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Option1,
		&i.Option2,
		&i.Option3,
		&i.Option4,
		&i.CorrectNum,
		&i.Explanation,
		&i.UserID,
		&i.UserName,
		&i.UserAvatarUrl,
		&i.CreatedAt,
	)
	return i, err
}

const fetchQuizCounts = `-- name: FetchQuizCounts :one
SELECT COUNT(*) FROM quizzes
`

func (q *Queries) FetchQuizCounts(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, fetchQuizCounts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const fetchQuizCountsByUserID = `-- name: FetchQuizCountsByUserID :one
SELECT COUNT(*) FROM quizzes
WHERE
  user_id = ?
`

func (q *Queries) FetchQuizCountsByUserID(ctx context.Context, userID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, fetchQuizCountsByUserID, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const fetchQuizzesByUserID = `-- name: FetchQuizzesByUserID :many
SELECT
  quizzes.id,
  quizzes.content,
  quizzes.option_1,
  quizzes.option_2,
  quizzes.option_3,
  quizzes.option_4,
  quizzes.correct_num,
  quizzes.explanation,
  quizzes.user_id,
  users.name AS user_name,
  users.avatar_url AS user_avatar_url,
  quizzes.created_at
FROM
  quizzes
INNER JOIN users ON quizzes.user_id = users.id
WHERE
  user_id = ?
LIMIT ? OFFSET ?
`

type FetchQuizzesByUserIDParams struct {
	UserID string `json:"user_id"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type FetchQuizzesByUserIDRow struct {
	ID            string         `json:"id"`
	Content       string         `json:"content"`
	Option1       string         `json:"option_1"`
	Option2       string         `json:"option_2"`
	Option3       sql.NullString `json:"option_3"`
	Option4       sql.NullString `json:"option_4"`
	CorrectNum    int8           `json:"correct_num"`
	Explanation   sql.NullString `json:"explanation"`
	UserID        string         `json:"user_id"`
	UserName      string         `json:"user_name"`
	UserAvatarUrl string         `json:"user_avatar_url"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (q *Queries) FetchQuizzesByUserID(ctx context.Context, arg FetchQuizzesByUserIDParams) ([]FetchQuizzesByUserIDRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchQuizzesByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchQuizzesByUserIDRow{}
	for rows.Next() {
		var i FetchQuizzesByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Option1,
			&i.Option2,
			&i.Option3,
			&i.Option4,
			&i.CorrectNum,
			&i.Explanation,
			&i.UserID,
			&i.UserName,
			&i.UserAvatarUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fetchRandomQuizzes = `-- name: FetchRandomQuizzes :many
SELECT
  quizzes.id,
  quizzes.content,
  quizzes.option_1,
  quizzes.option_2,
  quizzes.option_3,
  quizzes.option_4,
  quizzes.correct_num,
  quizzes.explanation,
  quizzes.user_id,
  users.name AS user_name,
  users.avatar_url AS user_avatar_url,
  quizzes.created_at
FROM
  quizzes
INNER JOIN users ON quizzes.user_id = users.id
ORDER BY RAND()
LIMIT ?
`

type FetchRandomQuizzesRow struct {
	ID            string         `json:"id"`
	Content       string         `json:"content"`
	Option1       string         `json:"option_1"`
	Option2       string         `json:"option_2"`
	Option3       sql.NullString `json:"option_3"`
	Option4       sql.NullString `json:"option_4"`
	CorrectNum    int8           `json:"correct_num"`
	Explanation   sql.NullString `json:"explanation"`
	UserID        string         `json:"user_id"`
	UserName      string         `json:"user_name"`
	UserAvatarUrl string         `json:"user_avatar_url"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (q *Queries) FetchRandomQuizzes(ctx context.Context, limit int32) ([]FetchRandomQuizzesRow, error) {
	rows, err := q.db.QueryContext(ctx, fetchRandomQuizzes, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FetchRandomQuizzesRow{}
	for rows.Next() {
		var i FetchRandomQuizzesRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Option1,
			&i.Option2,
			&i.Option3,
			&i.Option4,
			&i.CorrectNum,
			&i.Explanation,
			&i.UserID,
			&i.UserName,
			&i.UserAvatarUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertQuiz = `-- name: InsertQuiz :exec
INSERT INTO
  quizzes (id, user_id, content, option_1, option_2, option_3, option_4, correct_num, explanation, created_at)
VALUES
  (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		NOW()
  )
`

type InsertQuizParams struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Content     string         `json:"content"`
	Option1     string         `json:"option_1"`
	Option2     string         `json:"option_2"`
	Option3     sql.NullString `json:"option_3"`
	Option4     sql.NullString `json:"option_4"`
	CorrectNum  int8           `json:"correct_num"`
	Explanation sql.NullString `json:"explanation"`
}

func (q *Queries) InsertQuiz(ctx context.Context, arg InsertQuizParams) error {
	_, err := q.db.ExecContext(ctx, insertQuiz,
		arg.ID,
		arg.UserID,
		arg.Content,
		arg.Option1,
		arg.Option2,
		arg.Option3,
		arg.Option4,
		arg.CorrectNum,
		arg.Explanation,
	)
	return err
}
