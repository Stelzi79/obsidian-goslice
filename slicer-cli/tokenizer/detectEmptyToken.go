package tokenizer

func (EmptyLineToken) isToken() {}

type EmptyLineToken struct{}

func detectAndProcessEmptyToken(token string, tokens *Tokens, fmToken *FrontMatterToken, hdToken *Header3Token) bool {
	if detectedEmptyToken(token) {
		processEmptyLineToken(token, tokens)
		return true
	}
	if *fmToken != nil || *hdToken != nil {
		return false
	}
	return false
}

func detectedEmptyToken(token string) bool {
	return token == ""
}

func processEmptyLineToken(_ string, tokens *Tokens) {
	tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(EmptyLineToken{}))
}
