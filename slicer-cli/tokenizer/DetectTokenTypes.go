package tokenizer

func DetectTokenTypes(tokens *Tokens) {

	var fmToken FrontMatterToken
	var h3Token Header3Token
	fmOnlyOnce := false
	for _, token := range tokens.RawTokenList {
		// fmt.Println(token)

		// Check if the token starts with ---
		if detectAndProcessFrontMatterToken(string(token), tokens, &fmToken, &fmOnlyOnce) {
			continue
		}

		// Check if the token starts with ###
		if detectAndProcessHeaderToken(string(token), tokens, &h3Token) {
			continue
		}

		// Check and process empty line
		if detectAndProcessEmptyToken(string(token), tokens, &fmToken, &h3Token) {
			continue
		}

		// finally add unDetectedToken
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(UnDetectedToken(token)))

	}
}
