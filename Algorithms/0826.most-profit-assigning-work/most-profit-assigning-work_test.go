package problem0826

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	difficulty []int
	profit     []int
	worker     []int
	ans        int
}{

	{
		[]int{2, 4, 6, 8, 10},
		[]int{50, 20, 30, 40, 50},
		[]int{4, 5, 6, 7},
		200,
	},

	{
		[]int{2, 4, 6, 8, 10},
		[]int{10, 20, 30, 40, 50},
		[]int{4, 5, 6, 7},
		100,
	},

	// 可以有多个 testcase
}

func Test_maxProfitAssignment(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProfitAssignment(tc.difficulty, tc.profit, tc.worker), "输入:%v", tc)
	}
}

func Benchmark_maxProfitAssignment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProfitAssignment(tc.difficulty, tc.profit, tc.worker)
		}
	}
}
