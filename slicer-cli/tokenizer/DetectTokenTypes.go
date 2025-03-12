package tokenizer

func DetectTokenTypes(tokens *Tokens) {

	var fmToken FrontMatterToken
	for _, token := range tokens.RawTokenList {
		// fmt.Println(token)

		// Check and process empty line
		if detectAndProcessEmptyToken(string(token), tokens) {
			continue
		}

		// Check if the token starts with ---
		if detectAndProcessFrontMatterToken(string(token), tokens, &fmToken) {
			continue
		}

		// finally add unDetectedToken
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(UnDetectedToken{}))

	}
}
