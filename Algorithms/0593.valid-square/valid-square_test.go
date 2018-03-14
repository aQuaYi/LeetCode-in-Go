package problem0593

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	p1, p2, p3, p4 []int
	ans            bool
}{

	{
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 1},
		false,
	},

	{
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		false,
	},

	{
		[]int{1, 3},
		[]int{3, 7},
		[]int{7, 4},
		[]int{4, 0},
		false,
	},

	{
		[]int{0, 3},
		[]int{3, 7},
		[]int{7, 4},
		[]int{4, 0},
		true,
	},

	{
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
		[]int{0, 1},
		true,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, validSquare(tc.p1, tc.p2, tc.p3, tc.p4), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			validSquare(tc.p1, tc.p2, tc.p3, tc.p4)
		}
	}
}
