package repository

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"strings"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/repository/sqlite"
)

type SQLiteTokenRepository struct {
	cipher  cipher.Block
	queries *sqlite.Queries
}

func NewSQLiteTokenRepository(queries *sqlite.Queries) (*SQLiteTokenRepository, error) {
	cipher, err := aes.NewCipher([]byte(tokenEncryptionKey))
	if err != nil {
		return nil, err
	}

	return &SQLiteTokenRepository{
		cipher:  cipher,
		queries: queries,
	}, nil
}

func (r *SQLiteTokenRepository) Find(ctx context.Context, tokenStr string) (*tokens.Token, error) {
	segments := strings.SplitN(tokenStr, ":", 2)
	if len(segments) != 2 {
		return nil, tokens.ErrTokenNotFound
	}

	id := segments[1]
	token, err := r.queries.FindToken(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, tokens.ErrTokenNotFound
		}

		return nil, err
	}

	rawData, err := decrypt(r.cipher, token.Data)
	if err != nil {
		return nil, tokens.ErrUnableToDecrypt
	}

	return tokens.New(
		token.ID,
		tokens.WithData(rawData),
		tokens.WithVersion(token.Version),
	), nil
}

func (r *SQLiteTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	encryptedData, err := encrypt(r.cipher, token.Data())
	if err != nil {
		return tokens.ErrUnableToEncrypt
	}

	_, err = r.queries.CreateToken(ctx, sqlite.CreateTokenParams{
		ID:      token.Id(),
		Data:    encryptedData,
		Version: token.Version(),
	})

	return err
}
