package tokenizer

import (
	"bufio"
	"fmt"
	"os"
)

// LoadRawLinesTokensFromFile reads a file and puts the content into RawLines-Tokens
func LoadRawLinesTokensFromFile(filePath string, tokens *[]Token) {
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
