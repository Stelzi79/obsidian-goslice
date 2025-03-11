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

func ReadFile(filePath string, tokens *Tokens) {
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
