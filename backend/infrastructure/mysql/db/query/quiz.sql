-- name: FetchQuizByID :one
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
  quizzes.id = ?;

-- name: FetchQuizzesByUserID :many
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
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: FetchLatestQuizzes :many
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
LIMIT ? OFFSET ?;

-- name: FetchRandomQuizzes :many
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
LIMIT ?;

-- name: FetchQuizCounts :one
SELECT COUNT(*) FROM quizzes;

-- name: FetchQuizCountsByUserID :one
SELECT COUNT(*) FROM quizzes
WHERE
  user_id = ?;

-- name: InsertQuiz :exec
INSERT INTO
  quizzes (id, user_id, content, option_1, option_2, option_3, option_4, correct_num, explanation, created_at)
VALUES
  (
		sqlc.arg(id),
		sqlc.arg(user_id),
		sqlc.arg(content),
		sqlc.arg(option_1),
		sqlc.arg(option_2),
		sqlc.narg(option_3),
		sqlc.narg(option_4),
		sqlc.arg(correct_num),
		sqlc.narg(explanation),
		NOW()
  );

-- name: DeleteQuiz :exec
DELETE FROM quizzes
WHERE
  id = ?;
