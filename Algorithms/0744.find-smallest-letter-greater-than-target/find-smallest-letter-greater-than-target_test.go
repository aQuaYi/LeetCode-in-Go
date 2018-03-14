package problem0744

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	letters []byte
	target  byte
	ans     byte
}{

	{
		[]byte{'c', 'f', 'j'},
		'a',
		'c',
	},

	{
		[]byte{'c', 'f', 'j'},
		'c',
		'f',
	},

	{
		[]byte{'c', 'f', 'j'},
		'd',
		'f',
	},

	{
		[]byte{'c', 'f', 'j'},
		'g',
		'j',
	},

	{
		[]byte{'c', 'f', 'j'},
		'j',
		'c',
	},

	{
		[]byte{'c', 'f', 'j'},
		'k',
		'c',
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, nextGreatestLetter(tc.letters, tc.target), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextGreatestLetter(tc.letters, tc.target)
		}
	}
}
