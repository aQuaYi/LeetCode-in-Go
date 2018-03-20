package problem0781

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	answers []int
	ans     int
}{

	{
		[]int{1, 0, 1, 0, 0},
		5,
	},

	{
		[]int{1, 1, 2},
		5,
	},

	{
		[]int{10, 10, 10},
		11,
	},

	{
		[]int{},
		0,
	},

	// 可以有多个 testcase
}

func Test_numRabbits(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numRabbits(tc.answers), "输入:%v", tc)
	}
}

func Benchmark_numRabbits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numRabbits(tc.answers)
		}
	}
}
