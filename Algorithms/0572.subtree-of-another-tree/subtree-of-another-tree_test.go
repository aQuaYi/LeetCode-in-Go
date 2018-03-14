package problem0572

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	spre, sin []int
	tpre, tin []int
	ans       bool
}{

	{
		[]int{1, 1},
		[]int{1, 1},
		[]int{1},
		[]int{1},
		true,
	},

	{
		[]int{3, 4, 1, 2, 5},
		[]int{1, 4, 2, 3, 5},
		[]int{4, 1, 2},
		[]int{1, 4, 2},
		true,
	},

	{
		[]int{3, 4, 1, 5, 2},
		[]int{1, 4, 3, 2, 5},
		[]int{3, 1, 2},
		[]int{1, 3, 2},
		false,
	},

	{
		[]int{3, 4, 1, 2, 0, 5},
		[]int{1, 4, 0, 2, 3, 5},
		[]int{4, 1, 2},
		[]int{1, 4, 2},
		false,
	},

	// 可以有多个 testcase
}

func Test_isSubtree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		s := kit.PreIn2Tree(tc.spre, tc.sin)
		t := kit.PreIn2Tree(tc.tpre, tc.tin)
		ast.Equal(tc.ans, isSubtree(s, t), "输入:%v", tc)
	}
}

func Benchmark_isSubtree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			s := kit.PreIn2Tree(tc.spre, tc.sin)
			t := kit.PreIn2Tree(tc.tpre, tc.tin)
			isSubtree(s, t)
		}
	}
}
