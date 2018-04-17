package problem0806

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	widths []int
	S      string
	ans    []int
}{

	{
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		"",
		[]int{0, 0},
	},

	{
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		"abcdefghijklmnopqrstuvwxyz",
		[]int{3, 60},
	},

	{
		[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		"bbbcccdddaaa",
		[]int{2, 4},
	},

	// 可以有多个 testcase
}

func Test_numberOfLines(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numberOfLines(tc.widths, tc.S), "输入:%v", tc)
	}
}

func Benchmark_numberOfLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numberOfLines(tc.widths, tc.S)
		}
	}
}
