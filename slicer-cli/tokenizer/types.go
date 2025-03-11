package tokenizer

type Token struct {
	RawLine string
}

type Tokens struct {
	FilePath           string
	MoveToFolder       string
	RawTokenList       []Token
	ProcessedTokenList []Token
}
