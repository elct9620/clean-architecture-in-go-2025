package repository

import (
	"context"

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

func (r *InMemoryTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	r.tokens[token.Id()] = InMemoryTokenSchema{
		Id:      token.Id(),
		Data:    token.Data(),
		Version: token.Version(),
	}

	return nil
}
