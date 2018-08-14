package problem0874

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	commands  []int
	obstacles [][]int
	ans       int
}{

	{
		[]int{4, -1, 3},
		[][]int{},
		25,
	},

	{
		[]int{4, -1, 4, -2, 4},
		[][]int{{2, 4}},
		65,
	},

	// 可以有多个 testcase
}

func Test_robotSim(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, robotSim(tc.commands, tc.obstacles), "输入:%v", tc)
	}
}

func Benchmark_robotSim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			robotSim(tc.commands, tc.obstacles)
		}
	}
}
