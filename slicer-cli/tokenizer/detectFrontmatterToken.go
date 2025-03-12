package tokenizer

import (
	"strings"
)

func detectAndProcessFrontMatterToken(token string, tokens *Tokens, fmToken *FrontMatterToken) bool {

	// fmt.Println(token)
	if len(token) < 3 {
		return false
	}
	if token[0:3] == "---" && *fmToken == nil {
		// start of frontmatter
		fmToken = new(FrontMatterToken)
		return true
	}
	if token[0:3] == "---" && *fmToken != nil {
		// end of frontmatter
		if *fmToken == nil {
			panic("fmToken should be nil")
		}
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(*fmToken))
		*fmToken = nil
		return true
	}
	if fmToken != nil && strings.Contains(token, ":") {
		values := strings.SplitN(token, ":", 2)
		fmToken.Add(strings.TrimSpace(values[0]), strings.TrimSpace(values[1]))
	}

	return true

}
