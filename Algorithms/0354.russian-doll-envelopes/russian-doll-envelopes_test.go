package problem0354

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	envelopes [][]int
	ans       int
}{

	{
		[][]int{
			[]int{46, 89},
			[]int{50, 53},
			[]int{52, 68},
			[]int{72, 45},
			[]int{77, 81},
		},
		3,
	},

	{
		[][]int{
			[]int{5, 4},
			[]int{6, 4},
			[]int{6, 7},
			[]int{2, 3},
			[]int{100, 4},
		},
		3,
	},

	{
		[][]int{},
		0,
	},

	// 可以有多个 testcase
}

func Test_maxEnvelopes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxEnvelopes(tc.envelopes), "输入:%v", tc)
	}
}

func Benchmark_maxEnvelopes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxEnvelopes(tc.envelopes)
		}
	}
}
