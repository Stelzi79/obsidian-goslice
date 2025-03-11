package tokenizer

type Token struct {
	RawLine string
}

type RawToken string

type Tokens struct {
	FilePath           string
	MoveToFolder       string
	RawTokenList       []RawToken
	ProcessedTokenList []Token
}
