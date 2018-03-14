package problem0756

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	bottom  string
	allowed []string
	ans     bool
}{

	{
		"BCD",
		[]string{"ACC", "ACB", "ABD", "DAA", "BDC", "BDB", "DBC", "BBD", "BBC", "DBD", "BCC", "CDD", "ABA", "BAB", "DDC", "CCD", "DDA", "CCA", "DDD"},
		true,
	},

	{
		"XYZ",
		[]string{"XYD", "YZE", "DEA", "FFF"},
		true,
	},

	{
		"XXYX",
		[]string{"XXX", "XXY", "XYX", "XYY", "YXZ"},
		false,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, pyramidTransition(tc.bottom, tc.allowed), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pyramidTransition(tc.bottom, tc.allowed)
		}
	}
}
