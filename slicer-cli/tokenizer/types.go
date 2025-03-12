package tokenizer

type Token interface {
	isToken()
}

func (UnDetectedToken) isToken() {}

type RawToken string

type UnDetectedToken string

type Tokens struct {
	FilePath           string
	MoveToFolder       string
	RawTokenList       []RawToken
	ProcessedTokenList []Token
}
