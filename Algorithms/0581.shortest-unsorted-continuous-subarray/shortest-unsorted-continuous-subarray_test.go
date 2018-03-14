package problem0581

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0581(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2, 7, 4, 5, 6, 3, 8, 9}},
			ans{5},
		},

		question{
			para{[]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}},
			ans{8},
		},

		question{
			para{[]int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9}},
			ans{0},
		},

		question{
			para{[]int{2, 2, 2, 2, 2, 3, 4, 6, 4, 8, 10, 9, 15}},
			ans{5},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findUnsortedSubarray(p.one), "输入:%v", p)
	}
}

func Benchmark_findUnsortedSubarray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findUnsortedSubarray([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	}
}
