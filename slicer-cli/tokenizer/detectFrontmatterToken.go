package tokenizer

func detectAndProcessFrontMatterToken(token string, tokens *Tokens, fmToken *FrontMatterToken) bool {
	if fmToken == nil && len(token) < 3 {
		return false
	}
	if token[0:3] == "---" && fmToken == nil {
		fmToken = new(FrontMatterToken)
		*fmToken = make(FrontMatterToken)
		return true
	}
	if fmToken != nil {
		(*fmToken)[token] = token
	}
	if token[0:3] == "---" && fmToken != nil {
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(*fmToken))
		*fmToken = nil
		return true
	}

	return true

}

// func processFrontMatterToken(token string, tokens *Tokens, fmToken *FrontMatterToken) {

// 	// if len(token) > 3 {
// 	// 	if token[0:3] == "---" && !detectedFrontMatterStart {
// 	// 		detectedFrontMatterStart = true
// 	// 	} else if token[0:3] == "---" && detectedFrontMatterStart {
// 	// 		detectedFrontMatterStart = false
// 	// 		// detectedFrontMatterEnd := true
// 	// 	}
// 	// }

// 	panic("unimplemented")
// }

// func detectedFrontMatterToken(token string, fmToken *FrontMatterToken) bool {
// 	// if fmToken == nil && len(token) < 3 {
// 	// 	return false
// 	// }
// 	// if token[0:3] == "---" {
// 	// }
// 	return false
// }
