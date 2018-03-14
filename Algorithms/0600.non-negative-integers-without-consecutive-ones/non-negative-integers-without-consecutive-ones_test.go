package problem0600

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

	{1000, 144},

	{8, 6},

	{5, 5},

	{4, 4},

	// 可以有多个 testcase
}

func Test_findIntegers(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findIntegers(tc.num), "输入:%v", tc)
	}
}

func Benchmark_findIntegers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findIntegers(tc.num)
		}
	}
}
