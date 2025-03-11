package main

import (
	"fmt"

	"github.com/Stelzi79/obsidian-goslice/slicer-cli/tokenizer"
	"github.com/sanity-io/litter"
)

func main() {
	base_path := "/mnt/Repositories/knowlage-base/Personal Knowlage Base/New Stuff"
	selected_file := "Software.md"
	selected_moveToFolder := base_path + "/Software"

	fmt.Println("📢Obsidian-Goslice/slicer-cli")

	// var tokens []string
	var tokens tokenizer.Tokens
	tokens.RawTokenList = []tokenizer.Token{}
	tokens.FilePath = base_path + "/" + selected_file
	tokens.MoveToFolder = selected_moveToFolder

	// read in the file and put the content into RawLines-Tokens
	tokenizer.LoadRawLinesTokensFromFile(base_path+"/"+selected_file, &tokens.RawTokenList)

	fmt.Println("TestPrint:")
	litter.Dump(tokenizer.SelectN(10, tokens))
	// litter.Dump(tokenizer.SelectN(5, tokens))

}
