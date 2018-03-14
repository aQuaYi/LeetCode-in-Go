package problem0451

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	{
		"tree",
		"eert",
	},

	{
		"cccaaa",
		"aaaccc",
	},

	{
		"Aabb",
		"bbAa",
	},

	// 可以有多个 testcase
}

func Test_frequencySort(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, frequencySort(tc.s), "输入:%v", tc)
	}
}

func Benchmark_frequencySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			frequencySort(tc.s)
		}
	}
}
