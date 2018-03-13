package Problem0783

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	// TODO: 按照题目的格式转换 整型切片 到 TreeNode
	root []int

	ans int
}{

// 可以有多个 testcase
}

func Test_minDiffInBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minDiffInBST(root*TreeNode)(tc.J, tc.S), "输入:%v", tc)
	}
}

func Benchmark_minDiffInBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDiffInBST(root*TreeNode)(tc.J, tc.S)
		}
	}
}
