package problem0338

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans []int
}{

	{
		5,
		[]int{0, 1, 1, 2, 1, 2},
	},

	// 可以有多个 testcase
}

func Test_countBits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countBits(tc.num), "输入:%v", tc)
	}
}

func Benchmark_countBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countBits(1000)
	}
}
