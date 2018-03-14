package problem0135

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ratings []int
	ans     int
}{

	{
		[]int{1, 2, 3, 3, 3, 1},
		10,
	},

	{
		[]int{1, 2, 3, 3, 1},
		9,
	},

	{
		[]int{1, 2, 3, 2, 1},
		9,
	},

	{
		[]int{1, 2, 3, 4, 3},
		11,
	},

	{
		[]int{1, 2, 3, 4, 4},
		11,
	},

	{
		[]int{1, 2, 3, 4, 5},
		15,
	},

	{
		[]int{0},
		1,
	},

	// 可以有多个 testcase
}

func Test_candy(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, candy(tc.ratings), "输入:%v", tc)
	}
}

func Benchmark_candy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			candy(tc.ratings)
		}
	}
}
