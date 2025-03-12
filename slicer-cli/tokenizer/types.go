package tokenizer

type Token interface {
	isToken()
}

func (EmptyLineToken) isToken() {}
func (UnDetectedToken) isToken() {}

type RawToken string

type EmptyLineToken struct{}

type UnDetectedToken struct{}

type Tokens struct {
	FilePath           string
	MoveToFolder       string
	RawTokenList       []RawToken
	ProcessedTokenList []Token
}
