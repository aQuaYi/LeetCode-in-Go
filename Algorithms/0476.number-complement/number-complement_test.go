package problem0476

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num int
	ans int
}{

	{
		1,
		0,
	},

	{
		5,
		2,
	},

	// 可以有多个 testcase
}

func Test_findComplement(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findComplement(tc.num), "输入:%v", tc)
	}
}

func Benchmark_findComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findComplement(tc.num)
		}
	}
}
