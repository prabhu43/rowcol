package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rowcol",
	Short: "Filter rows and columns of tabular output of executing some cli commands",
	Long:  "Filter rows and columns of tabular output of executing some cli commands",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		if info.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("The command is intended to work with pipes.")
			return
		}

		if len(args) > 2 {
			fmt.Println("Invalid arguments! Check Help!")
			return
		}

		var wordsTable [][]string
		reader := bufio.NewReader(os.Stdin)

		for {
			line, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				break
			}
			lineWords := strings.Fields(string(line))
			wordsTable = append(wordsTable, lineWords)
		}

		filter, err := NewFilter(wordsTable, args[0], args[1])
		if err != nil {
			fmt.Println(err)
		} else {
			filter.Print()
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
