package repository

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"strings"

	"github.com/elct9620/clean-architecture-in-go-2025/internal/entity/tokens"
)

var tokenEncryptionKey = "0123456789abcdef" // AES-128

type InMemoryTokenSchema struct {
	Id      string
	Data    []byte
	Version string
}

type InMemoryTokenRepository struct {
	cipher cipher.Block
	tokens map[string]InMemoryTokenSchema
}

func NewInMemoryTokenRepository() (*InMemoryTokenRepository, error) {
	cipher, err := aes.NewCipher([]byte(tokenEncryptionKey))
	if err != nil {
		return nil, err
	}

	return &InMemoryTokenRepository{
		cipher: cipher,
		tokens: map[string]InMemoryTokenSchema{},
	}, nil
}

func (r *InMemoryTokenRepository) Find(ctx context.Context, tokenStr string) (*tokens.Token, error) {
	segments := strings.SplitN(tokenStr, ":", 2)
	if len(segments) != 2 {
		return nil, tokens.ErrTokenNotFound
	}

	id := segments[1]
	token, ok := r.tokens[id]
	if !ok {
		return nil, tokens.ErrTokenNotFound
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

func (r *InMemoryTokenRepository) Save(ctx context.Context, token *tokens.Token) error {
	encrypted, err := encrypt(r.cipher, token.Data())
	if err != nil {
		return tokens.ErrUnableToEncrypt
	}

	r.tokens[token.Id()] = InMemoryTokenSchema{
		Id:      token.Id(),
		Data:    encrypted,
		Version: token.Version(),
	}

	return nil
}

func encrypt(block cipher.Block, data []byte) ([]byte, error) {
	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)

	return encrypted, nil
}

func decrypt(block cipher.Block, data []byte) ([]byte, error) {
	iv := data[:aes.BlockSize]
	rawData := data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(rawData, rawData)

	return rawData, nil
}
