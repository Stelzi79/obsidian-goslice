package tokenizer

import "strings"

func (Header3Token) isToken() {}

type Header3Token []string

func detectAndProcessHeaderToken(token string, tokens *Tokens, h3Token *Header3Token) bool {

	if strings.HasPrefix(token, "###") && *h3Token != nil {
		// get the last token from the h3Token
		lastToken := (*h3Token)[len(*h3Token)-1]
		if lastToken == "" {
			// remove the last token
			*h3Token = (*h3Token)[:len(*h3Token)-1]
			// append an empty token
			tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(EmptyLineToken{}))
		}
		tokens.ProcessedTokenList = append(tokens.ProcessedTokenList, Token(*h3Token))
		*h3Token = nil
	}
	if strings.HasPrefix(token, "###") && *h3Token == nil {
		*h3Token = append(*h3Token, token)
		return true
	}
	if *h3Token != nil {
		*h3Token = append(*h3Token, token)
		return true
	}
	return false
}
