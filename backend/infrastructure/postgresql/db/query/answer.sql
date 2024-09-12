-- name: FetchAnswerCountsByQuizID :many
SELECT
  answer_num, 
  COUNT(*) AS count
FROM
  answers
WHERE
  quiz_id = $1
GROUP BY 
answer_num;

-- name: InsertAnswer :exec
INSERT INTO
  answers (quiz_id, user_id, answer_num, created_at)
VALUES
  (
    sqlc.arg(quiz_id),
    sqlc.arg(user_id),
    sqlc.arg(answer_num),
    NOW()
  );
