package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var srcRoot = ""
var files = []string{}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Go to Markdown (go2md) converts a .go file to markdown\n")
		fmt.Printf("Usage:\n\n")
		fmt.Printf("    go2md <file_1> <file_2> ... <file_n>\n\n")
		fmt.Printf("Optional Flags:\n\n")
		fmt.Printf("    -r <SRC_ROOT>n")
		os.Exit(0)
	}
	skip := false
	for index, arg := range os.Args[1:] {
		if arg == "-r" {
			if index >= len(os.Args)-2 {
				fmt.Printf("No arguments after -r\n")
				os.Exit(1)
			}
			srcRoot = os.Args[index+2]
			skip = true
		} else if !skip {
			files = append(files, arg)
		} else {
			skip = false
		}
	}

	for _, fileName := range files {
		os.Stdout.WriteString(file2md(fileName))
	}

}

func file2md(fileName string) string {

	os.Stderr.WriteString(fmt.Sprintf("%v\n", fileName))

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Printf("Unable to read file %s\n%v\n", fileName, err)
		os.Exit(1)
	}
	input := bufio.NewScanner(file)

	mdString := ""
	insideCodeBlock := false
	insideTestBlock := false
	ignoring := false
	lineCounter := 0

	const OpenCodeBlock = 0b10
	const CloseCodeBlock = 0b100
	const IncludeMarkdown = 0b1000
	const IncludeNormalLine = 0b10000
	const IncludeTestLine = 0b100000
	const IncludeLineFeed = 0b1000000

	for input.Scan() {

		lineCounter++
		var line = input.Text()
		var action = 0
		switch {
		case strings.HasPrefix(line, "// Ignore-On"):
			ignoring = true
			if insideCodeBlock {
				action = action | CloseCodeBlock
			}
		case strings.HasPrefix(line, "// Ignore-Off"):
			ignoring = false
			if insideCodeBlock {
				action = action | CloseCodeBlock
			}
		case strings.HasPrefix(line, "// "):
			ignoring = false
			action = action | IncludeMarkdown
			if insideCodeBlock {
				insideCodeBlock = false
				action = action | CloseCodeBlock
			}
		case strings.HasPrefix(line, "func Test_") && !ignoring:
			insideTestBlock = true
		case strings.HasPrefix(line, "}") && !ignoring && insideTestBlock:
			insideTestBlock = false
		case strings.TrimSpace(line) != "" && !ignoring:
			if !insideCodeBlock {
				action = action | OpenCodeBlock
				insideCodeBlock = true
			}
			if !insideTestBlock {
				action = action | IncludeNormalLine
			} else {
				action = action | IncludeTestLine
			}
		case !ignoring:
			action = action | IncludeLineFeed
		}
		if OpenCodeBlock == action&OpenCodeBlock {
			_, justFile := filepath.Split(fileName)
			mdString = mdString + fmt.Sprintf("\n\nSource: [%s](%s%s#L%d) | [Top](#top)\n\n", justFile, srcRoot, fileName, lineCounter)
			mdString = mdString + "\n``` go\n"
		}
		if CloseCodeBlock == action&CloseCodeBlock {
			mdString = mdString + "```\n"
		}
		if IncludeMarkdown == action&IncludeMarkdown {
			mdString = mdString + line[3:] + "\n"
		}
		if IncludeNormalLine == action&IncludeNormalLine {
			mdString = mdString + line + "\n"
		}
		if IncludeTestLine == action&IncludeTestLine {
			if strings.Contains(line, "assert.") {
				line = assert2equality(line)
			}
			mdString = mdString + line[1:] + "\n"
		}
		if IncludeLineFeed == action&IncludeLineFeed {
			mdString = mdString + "\n"
		}
	}
	if insideCodeBlock {
		mdString = mdString + "```\n"
	}
	return mdString
}

func assert2equality(originalLine string) string {
	line := strings.TrimRight(originalLine, " ")
	var equals bool = true
	if strings.Contains(line, "assert.Equal(t, ") {
		line = strings.Replace(line, "assert.Equal(t, ", "", 1)
	} else if strings.Contains(line, "assert.NotEqual(t, ") {
		line = strings.Replace(line, "assert.NotEqual(t, ", "", 1)
		equals = false
	}
	commaPosition := -1
	lastCommaPosition := -1
	pCounter := 0
	for inString, i := false, 0; i < len(line); i++ {
		if i > 0 && line[i] == byte('"') && line[i-1] != byte('\\') {
			inString = !inString
		}

		if !inString {
			if line[i] == byte('(') {
				pCounter++
			} else if line[i] == byte(')') {
				pCounter--
				lastCommaPosition = i
			} else if line[i] == byte(',') && pCounter == 0 {
				commaPosition = i
			}
		}

	}
	if commaPosition == -1 {
		panic("comma was supposed to be found")
	}

	a := line[0:commaPosition]
	b := line[commaPosition+2:]
	if pCounter != 0 && lastCommaPosition != -1 {
		b = line[commaPosition+2:lastCommaPosition] + line[lastCommaPosition+1:]
	}

	symbol := "⇔"
	if !equals {
		symbol = "⇎"
	}
	return fmt.Sprintf("%s %s %s", a, symbol, b)

}
