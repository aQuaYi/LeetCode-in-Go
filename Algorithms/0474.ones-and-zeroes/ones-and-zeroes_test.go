package problem0474

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	strs []string
	m    int
	n    int
	ans  int
}{

	{
		[]string{"111001", "10", "0001"},
		5,
		3,
		2,
	},

	{
		[]string{"1", "0", "10", "0001", "111001"},
		5,
		3,
		4,
	},

	{
		[]string{"10", "0001", "111001", "1", "0"},
		5,
		3,
		4,
	},

	{
		[]string{"10", "0", "1"},
		1,
		1,
		2,
	},

	// 可以有多个 testcase
}

func Test_findMaxForm(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findMaxForm(tc.strs, tc.m, tc.n), "输入:%v", tc)
	}
}

func Benchmark_findMaxForm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMaxForm(tc.strs, tc.m, tc.n)
		}
	}
}
