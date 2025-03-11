package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	base_path := "/mnt/Repositories/knowlage-base/Personal Knowlage Base/New Stuff"
	selected_file := "Software.md"
	selected_moveFolder := base_path + "/Software"

	num_lines := 1

	// open file handle to base_path
	fh, err := os.Open(base_path)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		return
	}
	defer fh.Close()
	fmt.Printf("🏷️Base path: '%s'\n", base_path)

	// read all files in base_path
	files, err := fh.Readdir(0)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		return
	}

	// for each file, read the content
	fmt.Printf("📁Files:\n")
	for _, file := range files {
		writtenBlocks := []Blocks{}
		if file.Name() == selected_file {
			fmt.Printf("\t📄 '%s'\n", file.Name())
			// open file handle to selected_file
			fhNewStuff, err := os.Open(base_path + "/" + selected_file)
			if err != nil {
				fmt.Printf("❌Error: %s\n", err)
				return
			}
			defer fhNewStuff.Close()
			fmt.Printf("\t🏷️opened: '%s'\n", selected_file)
			fmt.Printf("📄Content:\n")

			// read one line of the file
			blocks := readLines(fhNewStuff, num_lines, selected_moveFolder)
			printBlocks(blocks)

			// check if selected_moveFolder exists
			_, err = os.Stat(selected_moveFolder)
			if os.IsNotExist(err) {
				fmt.Printf("📁Creating folder '%s'\n", selected_moveFolder)
				err = os.Mkdir(selected_moveFolder, 0755)
				if err != nil {
					fmt.Printf("❌Error: %s\n", err)
					return
				}
			}
			// move the blocks to the selected_moveFolder
			writtenBlocks = append(writtenBlocks, writeNewNotes(blocks)...)

		}
		if len(writtenBlocks) == 0 {
			fmt.Printf("📄")
		} else {
			fmt.Printf("📄Blocks written:\n")
			// delete moved blocks from selected_file
			// fmt.Println(fh.Name())
			deleteBlocks(writtenBlocks, fh, file.Name())
		}

	}

}

func deleteBlocks(writtenBlocks []Blocks, fh *os.File, selected_file string) {
	newFileName := fh.Name() + "/new_" + selected_file
	fmt.Printf("📄New file: '%s'\n", newFileName)

	newFh, err := os.Create(newFileName)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		return
	}
	defer newFh.Close()
	fhSelected, err := os.Open(fh.Name() + "/" + selected_file)
	if err != nil {
		fmt.Printf("❌Error: %s\n", err)
		return
	}
	defer fhSelected.Close()
	fmt.Printf("📄Deleting blocks:\n")

	writtenLines := []string{}
	for _, block := range writtenBlocks {
		for _, line := range block.Lines {
			writtenLines = append(writtenLines, line)
		}
	}

	// scanner := bufio.NewScanner(fhSelected)
	// for scanner.Scan() {
	// 	line := scanner.Text()

	// 	if !contains(writtenLines, line) {
	// 		_, err := newFh.WriteString(line + "\n")
	// 		if err != nil {
	// 			fmt.Printf("❌Error: %s\n", err)
	// 			return
	// 		}
	// 	}
	// }

}

func writeNewNotes(blocks []Blocks) []Blocks {
	writtenBlocks := []Blocks{}
	for _, block := range blocks {
		fmt.Printf("📄Moving block '%s' to '%s'\n", block.MainTag, block.FileName)
		// open file handle to block.FileName
		fh, err := os.Create(block.FileName)
		if err != nil {
			fmt.Printf("❌Error: %s\n", err)
			break
		}
		defer fh.Close()
		fmt.Printf("\t🏷️opened: '%s'\n", block.FileName)
		writeProblems := false

		// write the block to the file
		for _, line := range block.Lines {
			_, err := fh.WriteString(line + "\n")
			if err != nil {
				fmt.Printf("❌Error: %s\n", err)
				writeProblems = true
				break
			}
		}
		if !writeProblems {
			writtenBlocks = append(writtenBlocks, block)
		}
	}
	return writtenBlocks
}

func printBlocks(blocks []Blocks) {
	fmt.Printf("📄Blocks:\n")
	for _, line := range blocks {
		fmt.Printf("\t📄MainTag '%s'\n\t📄FileName '%s'\n", line.MainTag, line.FileName)
		for _, subline := range line.Lines {
			fmt.Printf("\t\t📄 '%s'\n", subline)
		}

	}
}

type Blocks struct {
	MainTag  string
	FileName string
	Lines    []string
}

func readLines(fh *os.File, num_lines int, selected_moveFolder string) []Blocks {

	var returnedBlocks []Blocks
	returnedBlocks = nil
	scanner := bufio.NewScanner(fh)
	detectBlock := false
	for scanner.Scan() {
		var mainTag string
		// if Txt starts with "###" then start detecting block
		if strings.HasPrefix(scanner.Text(), "###") {
			if num_lines == 0 {
				break
			}
			headline := strings.TrimSpace(strings.TrimPrefix(scanner.Text(), "###"))
			headlines := strings.Split(headline, " ")
			mainTag = headlines[0]
			filename := selected_moveFolder + "/" + strings.TrimPrefix(mainTag, "#") + ".md"

			returnedBlocks = append(returnedBlocks, Blocks{MainTag: mainTag, FileName: filename, Lines: []string{}})

			detectBlock = true
			num_lines--
			// continue if line is empty
		} else if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if detectBlock {
			// append the line to the last block
			line := strings.TrimSpace(scanner.Text())
			returnedBlocks[len(returnedBlocks)-1].Lines = append(returnedBlocks[len(returnedBlocks)-1].Lines, line)
		}
	}
	return returnedBlocks
}
