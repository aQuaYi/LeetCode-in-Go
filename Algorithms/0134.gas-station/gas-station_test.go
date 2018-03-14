package problem0134

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	gas  []int
	cost []int
	ans  int
}{

	{
		[]int{4},
		[]int{5},
		-1,
	},

	{
		[]int{4, 3},
		[]int{5, 1},
		1,
	},
	// 可以有多个 testcase
}

func Test_canCompleteCircuit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, canCompleteCircuit(tc.gas, tc.cost), "输入:%v", tc)
	}
}

func Benchmark_canCompleteCircuit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canCompleteCircuit(tc.gas, tc.cost)
		}
	}
}
