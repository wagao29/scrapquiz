-- name: FindQuizByID :one
SELECT
   *
FROM
  quizzes
WHERE
  id = ?;

-- name: FindQuizzesByUserID :many
SELECT
   *
FROM
  quizzes
WHERE
  user_id = ?
LIMIT ? OFFSET ?;

-- name: FindLatestQuizzes :many
SELECT
  *
FROM
  quizzes
ORDER BY id DESC
LIMIT ? OFFSET ?;

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
