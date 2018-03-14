package problem0514

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ring string
	key  string
	ans  int
}{

	{
		"acbcccccbd",
		"abd",
		6,
	},

	{
		"godding",
		"gd",
		4,
	},

	// 可以有多个 testcase
}

func Test_findRotateSteps(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findRotateSteps(tc.ring, tc.key), "输入:%v", tc)
	}
}

func Benchmark_findRotateSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findRotateSteps(tc.ring, tc.key)
		}
	}
}
