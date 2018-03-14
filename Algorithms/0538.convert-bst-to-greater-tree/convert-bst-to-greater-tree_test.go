package problem0538

import (
	"fmt"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ansPost []int
}{

	{
		[]int{5, 2, 13},
		[]int{2, 5, 13},
		[]int{20, 13, 18},
	},

	// 可以有多个 testcase
}

func Test_convertBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		root := kit.PreIn2Tree(tc.pre, tc.in)
		ans := kit.Tree2Postorder(convertBST(root))
		ast.Equal(tc.ansPost, ans)
	}
}

func Benchmark_convertBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			root := kit.PreIn2Tree(tc.pre, tc.in)
			convertBST(root)
		}
	}
}
