package problem0617

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre1, in1 []int
	pre2, in2 []int
	ansPost   []int
}{

	{
		[]int{1, 3, 5, 2},
		[]int{5, 3, 1, 2},
		[]int{20, 10, 40, 30, 70},
		[]int{10, 40, 20, 30, 70},
		[]int{5, 40, 13, 70, 32, 21},
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		t1 := kit.PreIn2Tree(tc.pre1, tc.in1)
		t2 := kit.PreIn2Tree(tc.pre2, tc.in2)
		ansPost := kit.Tree2Postorder(mergeTrees(t1, t2))
		ast.Equal(tc.ansPost, ansPost, "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			t1 := kit.PreIn2Tree(tc.pre1, tc.in1)
			t2 := kit.PreIn2Tree(tc.pre2, tc.in2)
			mergeTrees(t1, t2)
		}
	}
}
