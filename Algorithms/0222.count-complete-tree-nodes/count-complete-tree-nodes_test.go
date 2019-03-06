package problem0222

import (
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pre, in []int
	count   int
}{

	{
		[]int{},
		[]int{},
		0,
	},

	{
		[]int{1, 2, 4, 5, 3, 6},
		[]int{4, 2, 5, 1, 6, 3},
		6,
	},
}

func Test_countNodes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		root := kit.PreIn2Tree(q.pre, q.in)
		ast.Equal(q.count, countNodes(root))
	}
}

func Benchmark_countNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			root := kit.PreIn2Tree(q.pre, q.in)
			countNodes(root)
		}
	}
}
