package problem0810

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{1, 1, 2},
		false,
	},

	// 可以有多个 testcase
}

func Test_xorGame(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, xorGame(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_xorGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			xorGame(tc.nums)
		}
	}
}
