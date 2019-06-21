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

func Test_Case(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd", "f", "kl"}
	streamChecker := Constructor(words)
	//
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

func Test_Case_2(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd"}
	streamChecker := Constructor(words)
	//
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
	ast.False(streamChecker.Query('c')) // return false
}

func Test_Case_3(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd"}
	streamChecker := Constructor(words)
	//
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
	ast.False(streamChecker.Query('d')) // return false
}

func Test_Case_4(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd"}
	streamChecker := Constructor(words)
	//
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
}

func Test_Case_5(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd", "cdd"}
	streamChecker := Constructor(words)
	//
	ast.False(streamChecker.Query('c')) // return false
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
	ast.True(streamChecker.Query('d'))  // return true, because 'cd' is in the wordlist
}

func Test_Case_6(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"abaa", "abaab", "aabbb", "bab", "ab"}
	streamChecker := Constructor(words)

	querys := []byte{'a', 'a', 'b', 'b', 'b', 'a', 'a', 'b', 'b', 'a', 'a', 'a', 'a', 'b', 'a', 'b', 'b', 'b', 'a', 'b', 'b', 'b', 'a', 'a', 'a', 'a', 'a', 'b', 'a', 'b', 'b', 'b', 'a', 'a', 'b', 'b', 'b', 'a', 'b', 'a'}
	expects := []bool{false, false, true, false, true, false, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, false, false, true, false, true, false, false, false, false, true, false, true, false, true, false}

	ast.Equal(len(querys), len(expects))
	for i, q := range querys {
		ast.Equal(expects[i], streamChecker.Query(q), "%d:%s", i, string(querys[:i+1]))
	}
}

func Test_Case_7(t *testing.T) {
	ast := assert.New(t)
	//
	words := []string{"cd"}
	streamChecker := Constructor(words)
	//
	ast.False(streamChecker.Query('c')) // return false
	ast.False(streamChecker.Query('a')) // return false
	ast.False(streamChecker.Query('d')) // return false
}
