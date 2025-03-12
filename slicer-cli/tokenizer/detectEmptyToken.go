package tokenizer

func detectAndProcessEmptyToken(token string, tokens *Tokens) bool {

	if detectedEmptyToken(token) {
		processEmptyLineToken(token, tokens)
		return true
	} else {
		return false
	}
}

func detectedEmptyToken(token string) bool {
	return token == ""
}

func processEmptyLineToken(_ string, tokens *Tokens) {
	tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(EmptyLineToken{}))
}
