package repository

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"strings"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
	bolt "go.etcd.io/bbolt"
)

const BoltTokensTableName = "tokens"

type BoltTokenSchema struct {
	Id      string
	Data    []byte
	Version string
}

type BoltTokenRepository struct {
	cipher    cipher.Block
	db        *bolt.DB
	tableName string
}

func NewBoltTokenRepository(db *bolt.DB) (*BoltTokenRepository, error) {
	cipher, err := aes.NewCipher([]byte(tokenEncryptionKey))
	if err != nil {
		return nil, err
	}

	return &BoltTokenRepository{
		cipher:    cipher,
		db:        db,
		tableName: BoltTokensTableName,
	}, nil
}

func (r *BoltTokenRepository) Find(ctx context.Context, tokenStr string) (*tokens.Token, error) {
	segments := strings.SplitN(tokenStr, ":", 2)
	if len(segments) != 2 {
		return nil, tokens.ErrTokenNotFound
	}

	id := segments[1]
	var token BoltTokenSchema
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(r.tableName))
		if bucket == nil {
			return tokens.ErrTokenNotFound
		}

		boltValue := bucket.Get([]byte(id))
		if boltValue == nil {
			return tokens.ErrTokenNotFound
		}

		return json.Unmarshal(boltValue, &token)
	})

	if err != nil {
		return nil, err
	}

	rawData, err := decrypt(r.cipher, token.Data)
	if err != nil {
		return nil, tokens.ErrUnableToDecrypt
	}

	return tokens.New(
		token.Id,
		tokens.WithVersion(token.Version),
		tokens.WithData(rawData),
	), nil
}

func (r *BoltTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	encryptedData, err := encrypt(r.cipher, token.Data())
	if err != nil {
		return tokens.ErrUnableToEncrypt
	}

	return r.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(r.tableName))
		if err != nil {
			return err
		}

		tokenSchema := BoltTokenSchema{
			Id:      token.Id(),
			Data:    encryptedData,
			Version: token.Version(),
		}

		data, err := json.Marshal(tokenSchema)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(token.Id()), data)
	})
}
