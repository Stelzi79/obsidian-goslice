package tokenizer

import (
	"bufio"
	"fmt"
	"os"
)

type Token struct {
	RawLine string
}

type Tokens []Token

// LoadRawLinesTokensFromFile reads a file and puts the content into RawLines-Tokens
func LoadRawLinesTokensFromFile(filePath string, tokens *Tokens) {
	// open file
	fh, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		*tokens = append(*tokens, Token{RawLine: scanner.Text()})
	}
}

func SelectN(n int, tokens Tokens) Tokens {
	if n > len(tokens) {
		return tokens
	}
	return tokens[:n]
}
