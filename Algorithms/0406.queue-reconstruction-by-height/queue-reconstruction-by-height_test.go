package problem0406

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	people [][]int
	ans    [][]int
}{

	{
		[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
		[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
	},

	// 可以有多个 testcase
}

func Test_reconstructQueue(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, reconstructQueue(tc.people), "输入:%v", tc)
	}
}

func Benchmark_reconstructQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reconstructQueue(tc.people)
		}
	}
}
