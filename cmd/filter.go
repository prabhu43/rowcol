package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type Filter struct {
	wordsTable [][]string
	rows       Limit
	cols       Limit
}

type Limit struct {
	from int
	to   int
}

func (f *Filter) Print() {
	fmt.Println("printing output:", f.rows, f.cols)

	for i := f.rows.from - 1; i < f.rows.to; i++ {
		for j := f.cols.from - 1; j < len(f.wordsTable[i]) && j < f.cols.to; j++ {
			fmt.Printf("%s,", f.wordsTable[i][j])
		}
		fmt.Println()
	}
}

func NewFilter(wordsTable [][]string, rows, cols string) (*Filter, error) {

	rowsLimit, err := getLimit(rows, "Invalid rows")
	if err != nil {
		return nil, err
	}
	if rowsLimit.from > len(wordsTable) || rowsLimit.to > len(wordsTable) {
		return nil, fmt.Errorf("Invalid rows")
	}

	colsLimit, err := getLimit(cols, "Invalid cols")
	if err != nil {
		return nil, err
	}

	return &Filter{wordsTable, rowsLimit, colsLimit}, nil
}

func getLimit(limitsStr string, errMsg string) (Limit, error) {
	limits := strings.FieldsFunc(limitsStr, func(r rune) bool {
		return r == ':'
	})
	if len(limits) > 2 {
		return Limit{}, fmt.Errorf(errMsg)
	}

	var rowsLimit Limit
	switch len(limits) {
	case 1:
		fromRow, err := strconv.Atoi(limits[0])
		if err != nil {
			return Limit{}, err
		}
		rowsLimit = Limit{fromRow, fromRow}
	case 2:
		fromRow, err := strconv.Atoi(limits[0])
		if err != nil {
			return Limit{}, err
		}
		toRow, err := strconv.Atoi(limits[1])
		if err != nil {
			return Limit{}, err
		}
		rowsLimit = Limit{fromRow, toRow}

	default:
		return Limit{}, fmt.Errorf(errMsg)
	}
	if rowsLimit.from > rowsLimit.to {
		return Limit{}, fmt.Errorf(errMsg)
	}
	return rowsLimit, nil
}
