package problem0915

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	ans int
}{

	{
		[]int{5, 0, 3, 8, 6},
		3,
	},

	{
		[]int{1, 1, 1, 0, 6, 12},
		4,
	},

	// 可以有多个 testcase
}

func Test_partitionDisjoint(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, partitionDisjoint(tc.A), "输入:%v", tc)
	}
}

func Benchmark_partitionDisjoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			partitionDisjoint(tc.A)
		}
	}
}
