package Problem0127

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0127(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		beginWord string
		endWord   string
		wordList  []string
		ans       int
	}{

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			0,
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog", "dot"},
			3,
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			5,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, ladderLength(tc.beginWord, tc.endWord, tc.wordList), "输入:%v", tc)
	}
}
