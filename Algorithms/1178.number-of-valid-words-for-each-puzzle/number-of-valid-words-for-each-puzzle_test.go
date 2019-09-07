package problem1178

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	words []string
 puzzles []string
	ans []int
}{



	// 可以有多个 testcase
}

func Test_findNumOfValidWords(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findNumOfValidWords(tc.words, tc.puzzles), "输入:%v", tc)
	}
}

func Benchmark_findNumOfValidWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findNumOfValidWords(tc.words, tc.puzzles)
		}
	}
}
