package problem0881

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	people []int
	limit  int
	ans    int
}{

	{
		[]int{1, 2},
		3,
		1,
	},

	{
		[]int{3, 2, 2, 1},
		3,
		3,
	},

	{
		[]int{3, 5, 3, 4},
		5,
		4,
	},

	// 可以有多个 testcase
}

func Test_numRescueBoats(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, numRescueBoats(tc.people, tc.limit), "输入:%v", tc)
	}
}

func Benchmark_numRescueBoats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numRescueBoats(tc.people, tc.limit)
		}
	}
}
