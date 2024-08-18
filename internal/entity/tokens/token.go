package tokens

import "errors"

const (
	CurrentVersion = "v1"
)

var ErrTokenNotFound = errors.New("token not found")

type TokenOption func(*Token)

type Token struct {
	id      string
	data    []byte
	version string
}

func New(id string, opts ...TokenOption) *Token {
	token := &Token{
		id:      id,
		version: CurrentVersion,
	}

	for _, opt := range opts {
		opt(token)
	}

	return token
}

func (t *Token) Id() string {
	return t.id
}

func (t *Token) Data() []byte {
	return t.data
}

func (t *Token) Version() string {
	return t.version
}

func (t *Token) SetData(data []byte) {
	t.data = data
}

func (t *Token) Raw() string {
	return string(t.data)
}

func (t Token) String() string {
	return t.version + ":" + t.id
}

func WithVersion(version string) func(*Token) {
	return func(t *Token) {
		t.version = version
	}
}

func WithData(data []byte) func(*Token) {
	return func(t *Token) {
		t.data = data
	}
}
