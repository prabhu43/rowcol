package cmd_test

import (
	"testing"

	"github.com/prabhu43/rowcol/cmd"
	"github.com/stretchr/testify/assert"
)

type input struct {
	row string
	col string
}

type scenario struct {
	in        input
	outputRow cmd.Limit
	outputCol cmd.Limit
}

func TestNewFilter(t *testing.T) {
	wordsTable := make([][]string, 1)
	_, err := cmd.NewFilter(wordsTable, "1", "1")
	assert.NoError(t, err)
}

func TestNewFilterSuccessIfRowIsValid(t *testing.T) {
	scenarios := []scenario{
		{in: input{"1", "1"}, outputRow: cmd.NewLimit(1, 1), outputCol: cmd.NewLimit(1, 1)},
		{in: input{"1:1", "1"}, outputRow: cmd.NewLimit(1, 1), outputCol: cmd.NewLimit(1, 1)},
		{in: input{"1:2", "1"}, outputRow: cmd.NewLimit(1, 2), outputCol: cmd.NewLimit(1, 1)},
	}
	wordsTable := make([][]string, 2)

	for _, scenario := range scenarios {
		filter, err := cmd.NewFilter(wordsTable, scenario.in.row, scenario.in.col)
		assert.NoError(t, err)
		assert.Equal(t, scenario.outputRow, filter.Rows)
		assert.Equal(t, scenario.outputCol, filter.Cols)
	}
}

func TestNewFilterThrowErrorIfRowIsInvalid(t *testing.T) {
	scenarios := []input{
		{"qwe", "1"},
		{"@#@#", "1"},
		{"2", "1"},
		{"1:0", "1"},
	}
	wordsTable := make([][]string, 1)

	for _, limit := range scenarios {
		_, err := cmd.NewFilter(wordsTable, limit.row, limit.col)
		assert.Error(t, err)
	}
}

func TestNewFilterSuccessIfColumnIsValid(t *testing.T) {
	scenarios := []scenario{
		{in: input{"1", "1"}, outputRow: cmd.NewLimit(1, 1), outputCol: cmd.NewLimit(1, 1)},
		{in: input{"1", "1:1"}, outputRow: cmd.NewLimit(1, 1), outputCol: cmd.NewLimit(1, 1)},
		{in: input{"1", "1:2"}, outputRow: cmd.NewLimit(1, 1), outputCol: cmd.NewLimit(1, 2)},
	}
	wordsTable := make([][]string, 2)

	for _, scenario := range scenarios {
		filter, err := cmd.NewFilter(wordsTable, scenario.in.row, scenario.in.col)
		assert.NoError(t, err)
		assert.Equal(t, scenario.outputRow, filter.Rows)
		assert.Equal(t, scenario.outputCol, filter.Cols)
	}
}

func TestNewFilterThrowErrorIfColumnIsInvalid(t *testing.T) {
	scenarios := []input{
		{"1", "qwe"},
		{"1", "@#@#"},
		{"1", "1:0"},
	}
	wordsTable := make([][]string, 1)

	for _, limit := range scenarios {
		_, err := cmd.NewFilter(wordsTable, limit.row, limit.col)
		assert.Error(t, err)
	}
}
