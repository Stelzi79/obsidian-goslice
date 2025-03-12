package tokenizer

import (
	"strings"
)

func (FrontMatterToken) isToken() {}

type FrontMatterToken map[string]string

func (fmToken *FrontMatterToken) Add(key string, value string) {
	// fmt.Printf("Adding key: %s, value: %s\n", key, value)
	if *fmToken == nil {
		*fmToken = make(FrontMatterToken)
	}
	(*fmToken)[key] = value
}

func detectAndProcessFrontMatterToken(token string, tokens *Tokens, fmToken *FrontMatterToken, fmDetected *bool) bool {

	// fmt.Println(token)
	if len(token) < 3 || *fmDetected {
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
		*fmDetected = true
		return true
	}
	if fmToken != nil && !*fmDetected && strings.Contains(token, ":") {
		values := strings.SplitN(token, ":", 2)
		fmToken.Add(strings.TrimSpace(values[0]), strings.TrimSpace(values[1]))
		return true
	}
	return false
}
