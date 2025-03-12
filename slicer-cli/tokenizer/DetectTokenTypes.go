package tokenizer

func DetectTokenTypes(tokens *Tokens) {

	detectedFrontMatterStart := false
	// detectedFrontMatterEnd := false

	for _, token := range tokens.RawTokenList {
		// fmt.Println(token)

		// Check and process empty line
		if detectAndProcessEmptyToken(string(token), tokens) {
			continue
		}

		// Check if the token starts with ---
		if len(token) > 3 {
			if token[0:3] == "---" && !detectedFrontMatterStart {
				detectedFrontMatterStart = true
			} else if token[0:3] == "---" && detectedFrontMatterStart {
				detectedFrontMatterStart = false
				// detectedFrontMatterEnd := true
			}
		}

		// add unDetectedToken
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(UnDetectedToken{}))

	}
}
