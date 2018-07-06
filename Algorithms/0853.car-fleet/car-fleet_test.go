package problem0853

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target   int
	position []int
	speed    []int
	ans      int
}{

	{
		12,
		[]int{10, 8, 0, 5, 3},
		[]int{2, 4, 1, 1, 3},
		3,
	},

	// 可以有多个 testcase
}

func Test_carFleet(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, carFleet(tc.target, tc.position, tc.speed), "输入:%v", tc)
	}
}

func Benchmark_carFleet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			carFleet(tc.target, tc.position, tc.speed)
		}
	}
}
