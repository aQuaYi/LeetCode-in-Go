package problem0925

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	name  string
	typed string
	ans   bool
}{

	{
		"alex",
		"aaaaaaaaaaaallllllllllllllleeeeeeeeeeeeeeeexxxxxxxxx",
		true,
	},

	{
		"alex",
		"aal",
		false,
	},

	{
		"alex",
		"aaleexOK",
		false,
	},

	{
		"alex",
		"aaleex",
		true,
	},

	{
		"saeed",
		"ssaaedd",
		false,
	},

	{
		"leelee",
		"lleeelee",
		true,
	},

	{
		"laiden",
		"laiden",
		true,
	},

	// 可以有多个 testcase
}

func Test_isLongPressedName(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isLongPressedName(tc.name, tc.typed), "输入:%v", tc)
	}
}

func Benchmark_isLongPressedName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isLongPressedName(tc.name, tc.typed)
		}
	}
}
