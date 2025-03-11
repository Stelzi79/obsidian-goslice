package tokenizer

func SelectN(n int, tokens Tokens) Tokens {
	if n > len(tokens.RawTokenList) {
		return tokens
	}
	var newTokens = tokens
	newTokens.RawTokenList = make([]RawToken, n)
	copy(newTokens.RawTokenList, tokens.RawTokenList[:n])
	return newTokens
}
