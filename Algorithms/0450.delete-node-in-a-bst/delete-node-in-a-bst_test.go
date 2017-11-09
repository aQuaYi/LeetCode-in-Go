package Problem0450

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	key     int
	ans     []int
}{

	{
		[]int{5, 3, 2, 4, 6, 7},
		[]int{2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 2, 4, 6, 7},
	},

	// 可以有多个 testcase
}

func Test_deleteNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, deleteNode(tc.root, tc.key), "输入:%v", tc)
	}
}

func Benchmark_deleteNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deleteNode(tc.root, tc.key)
		}
	}
}
