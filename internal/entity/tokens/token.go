package tokens

const (
	CurrentVersion = "v1"
)

type Token struct {
	id      string
	data    []byte
	version string
}

func New(id string) *Token {
	return &Token{
		id:      id,
		version: CurrentVersion,
	}
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
