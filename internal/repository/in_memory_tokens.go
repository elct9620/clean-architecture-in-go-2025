package repository

import (
	"context"
	"strings"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
)

type InMemoryTokenSchema struct {
	Id      string
	Data    []byte
	Version string
}

type InMemoryTokenRepository struct {
	tokens map[string]InMemoryTokenSchema
}

func NewInMemoryTokenRepository() *InMemoryTokenRepository {
	return &InMemoryTokenRepository{
		tokens: map[string]InMemoryTokenSchema{},
	}
}

func (r *InMemoryTokenRepository) Find(ctx context.Context, id string) (*tokens.Token, error) {
	id = strings.SplitN(id, ":", 2)[1]
	token, ok := r.tokens[id]
	if !ok {
		return nil, tokens.ErrTokenNotFound
	}

	return tokens.New(
		token.Id,
		tokens.WithVersion(token.Version),
		tokens.WithData(token.Data),
	), nil
}

func (r *InMemoryTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	r.tokens[token.Id()] = InMemoryTokenSchema{
		Id:      token.Id(),
		Data:    token.Data(),
		Version: token.Version(),
	}

	return nil
}
