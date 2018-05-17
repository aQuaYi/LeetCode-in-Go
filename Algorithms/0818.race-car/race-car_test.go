package problem0818

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	target int
	ans    int
}{

	{
		3,
		2,
	},

	{
		6,
		5,
	},

	// 可以有多个 testcase
}

func Test_racecar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, racecar(tc.target), "输入:%v", tc)
	}
}

func Benchmark_racecar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			racecar(tc.target)
		}
	}
}
