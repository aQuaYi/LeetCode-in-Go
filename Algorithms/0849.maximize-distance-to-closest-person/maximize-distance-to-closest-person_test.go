package problem0849

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	seats []int
	ans   int
}{

	{
		[]int{1, 0, 0, 0, 1, 0, 1},
		2,
	},

	{
		[]int{0, 0, 0, 1},
		3,
	},

	{
		[]int{1, 0, 0, 0},
		3,
	},

	// 可以有多个 testcase
}

func Test_maxDistToClosest(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxDistToClosest(tc.seats), "输入:%v", tc)
	}
}

func Benchmark_maxDistToClosest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxDistToClosest(tc.seats)
		}
	}
}
