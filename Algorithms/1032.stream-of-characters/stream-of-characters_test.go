package problem1032

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
	ans   StreamChecker
}{

	// 可以有多个 testcase
}

func Test_Constructor(t *testing.T) {
	ast := assert.New(t)

	words := []string{"cd", "f", "kl"}
	streamChecker := Constructor(words)

	ast.False(streamChecker.Query('a')) // return false
	ast.False(streamChecker.Query('b')) // return false
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
	ast.False(streamChecker.Query('e')) // return false
	ast.True(streamChecker.Query('f'))  // return true, because 'f' is in the wordlist
	ast.False(streamChecker.Query('g')) // return false
	ast.False(streamChecker.Query('h')) // return false
	ast.False(streamChecker.Query('i')) // return false
	ast.False(streamChecker.Query('j')) // return false
	ast.False(streamChecker.Query('k')) // return false
	ast.True(streamChecker.Query('l'))  // return true, because 'kl' is in the wordlist
}
