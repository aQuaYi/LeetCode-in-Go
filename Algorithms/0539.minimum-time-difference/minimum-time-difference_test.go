package problem0539

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	timePoints []string
	ans        int
}{

	{
		[]string{"12:08", "13:15", "23:59"},
		67,
	},

	{
		[]string{"23:59", "00:00"},
		1,
	},

	// 可以有多个 testcase
}

func Test_findMinDifference(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMinDifference(tc.timePoints), "输入:%v", tc)
	}
}

func Benchmark_findMinDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMinDifference(tc.timePoints)
		}
	}
}
