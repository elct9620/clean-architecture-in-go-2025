-- name: FindToken :one
SELECT * FROM tokens WHERE id = ?
LIMIT 1;

-- name: CreateToken :one
INSERT INTO tokens (
  id, data, version
) VALUES (
  ?, ?, ?
)
RETURNING *;
