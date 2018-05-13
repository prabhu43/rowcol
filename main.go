package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		return
	}

	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	fmt.Println("printing output:")
	wordsTable := make([][]string, len(lines))

	for i := 0; i < len(lines); i++ {
		lineWords := strings.Fields(lines[i])
		wordsTable[i] = lineWords
	}

	for i := 0; i < len(wordsTable); i++ {
		for j := 0; j < len(wordsTable[i]); j++ {
			fmt.Printf("%s,", wordsTable[i][j])
		}
		fmt.Println()
	}

}
