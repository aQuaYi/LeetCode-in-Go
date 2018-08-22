package problem0886

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N        int
	dislikes [][]int
	ans      bool
}{

	{
		4,
		[][]int{{1, 2}, {1, 3}, {2, 4}},
		true,
	},

	{
		3,
		[][]int{{1, 2}, {1, 3}, {2, 3}},
		false,
	},

	{
		5,
		[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}},
		false,
	},

	// 可以有多个 testcase
}

func Test_possibleBipartition(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, possibleBipartition(tc.N, tc.dislikes), "输入:%v", tc)
	}
}

func Benchmark_possibleBipartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			possibleBipartition(tc.N, tc.dislikes)
		}
	}
}
