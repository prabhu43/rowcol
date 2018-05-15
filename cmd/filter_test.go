package cmd_test

import (
	"testing"

	"github.com/prabhu43/rowcol/cmd"
	"github.com/stretchr/testify/assert"
)

func TestNewFilter(t *testing.T) {
	wordsTable := make([][]string, 1)
	_, err := cmd.NewFilter(wordsTable, "1", "1")
	assert.NoError(t, err)
}

func TestNewFilterThrowErrorIfRowIsInvalid(t *testing.T) {
	wordsTable := make([][]string, 1)

	_, err := cmd.NewFilter(wordsTable, "qwe", "1")
	assert.Error(t, err)

	_, err = cmd.NewFilter(wordsTable, "@#@#", "1")
	assert.Error(t, err)

	_, err = cmd.NewFilter(wordsTable, "2", "1")
	assert.Error(t, err)

	_, err = cmd.NewFilter(wordsTable, "1:0", "1")
	assert.Error(t, err)
}

func TestNewFilterThrowErrorIfColIsInvalid(t *testing.T) {
	wordsTable := make([][]string, 1)

	_, err := cmd.NewFilter(wordsTable, "1", "qwe")
	assert.Error(t, err)

	_, err = cmd.NewFilter(wordsTable, "1", "@#@#")
	assert.Error(t, err)

	_, err = cmd.NewFilter(wordsTable, "1", "1:0")
	assert.Error(t, err)
}
