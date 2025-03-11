package tokenizer

import (
	"bufio"
	"fmt"
	"os"
)

// LoadRawLinesTokensFromFile reads a file and puts the content into RawLines-Tokens
func LoadRawLinesTokensFromFile(tokens *Tokens) {
	// open file

	fh, err := os.Open(tokens.FilePath)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		tokens.RawTokenList = append(tokens.RawTokenList, RawToken(scanner.Text()))
	}
}
