// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package sqlite

import (
	"context"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  id, data, version
) VALUES (
  ?, ?, ?
)
RETURNING id, data, version
`

type CreateTokenParams struct {
	ID      string
	Data    []byte
	Version string
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken, arg.ID, arg.Data, arg.Version)
	var i Token
	err := row.Scan(&i.ID, &i.Data, &i.Version)
	return i, err
}

const findToken = `-- name: FindToken :one
SELECT id, data, version FROM tokens WHERE id = ?
LIMIT 1
`

func (q *Queries) FindToken(ctx context.Context, id string) (Token, error) {
	row := q.db.QueryRowContext(ctx, findToken, id)
	var i Token
	err := row.Scan(&i.ID, &i.Data, &i.Version)
	return i, err
}
