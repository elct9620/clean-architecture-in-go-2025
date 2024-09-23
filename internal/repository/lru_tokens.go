package repository

import (
	"context"
	"log"
	"time"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
	"github.com/elct9620/clean-architecture-in-go-2025/internal/usecase"
	"github.com/hashicorp/golang-lru/v2/expirable"
)

const (
	DefaultCacheSize = 1000
	DefaultCacheTTL  = 60
)

type CachableTokenRepository usecase.TokenRepository

type LruTokenRepository struct {
	cache  *expirable.LRU[string, InMemoryTokenSchema]
	tokens CachableTokenRepository
}

func NewLruTokenRepository(tokens CachableTokenRepository) *LruTokenRepository {
	cache := expirable.NewLRU[string, InMemoryTokenSchema](DefaultCacheSize, nil, DefaultCacheTTL*time.Second)

	return &LruTokenRepository{
		cache:  cache,
		tokens: tokens,
	}
}

func (r *LruTokenRepository) Find(ctx context.Context, tokenStr string) (*tokens.Token, error) {
	token, ok := r.cache.Get(tokenStr)
	if ok {
		log.Printf("Cache hit: %s", tokenStr)

		entity := tokens.New(
			token.Id,
			tokens.WithVersion(token.Version),
			tokens.WithData(token.Data),
		)

		return entity, nil
	}

	log.Printf("Cache miss: %s", tokenStr)

	entity, err := r.tokens.Find(ctx, tokenStr)
	if err != nil {
		return nil, err
	}

	r.cache.Add(tokenStr, InMemoryTokenSchema{
		Id:      entity.Id(),
		Data:    entity.Data(),
		Version: entity.Version(),
	})

	return entity, nil
}

func (r *LruTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	return r.tokens.Save(ctx, token)
}
