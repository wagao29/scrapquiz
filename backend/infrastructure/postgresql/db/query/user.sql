-- name: UserFindAll :many
SELECT
   *
FROM
   users;

-- name: UserFindByID :one
SELECT
   *
FROM
   users
WHERE
   id = $1;

-- name: InsertUser :exec
INSERT INTO
   users (id, name, avatar_url, created_at, updated_at)
VALUES
   (
      sqlc.arg(id),
      sqlc.arg(name),
      sqlc.arg(avatar_url),
      NOW(),
      NOW()
   );

-- name: UpdateUser :exec
UPDATE users
SET
   name = sqlc.arg(name),
   avatar_url = sqlc.arg(avatar_url),
   updated_at = NOW()
WHERE
   id = sqlc.arg(id);

-- name: DeleteUser :exec
DELETE FROM users
WHERE
   id = $1;
