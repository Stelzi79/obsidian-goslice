package tokenizer

import "fmt"

type Token interface {
	isToken()
}

func (EmptyLineToken) isToken()   {}
func (UnDetectedToken) isToken()  {}
func (FrontMatterToken) isToken() {}

type RawToken string

type EmptyLineToken struct{}
type UnDetectedToken struct{}
type FrontMatterToken map[string]string

func (fmToken *FrontMatterToken) Add(key string, value string) {
	fmt.Printf("Adding key: %s, value: %s\n", key, value)
	if *fmToken == nil {
		*fmToken = make(FrontMatterToken)
	}
	(*fmToken)[key] = value
}

type Tokens struct {
	FilePath           string
	MoveToFolder       string
	RawTokenList       []RawToken
	ProcessedTokenList []Token
}
