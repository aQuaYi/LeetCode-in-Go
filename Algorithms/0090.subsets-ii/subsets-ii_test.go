package problem0090

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
type para struct {
	nums []int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0090(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 2},
			},
			ans{
				[][]int{
					[]int{},
					[]int{1},
					[]int{1, 2},
					[]int{1, 2, 2},
					[]int{2},
					[]int{2, 2},
				},
			},
		},

		question{
			para{
				[]int{1, 2, 2, 3},
			},
			ans{
				[][]int{
					[]int{},
					[]int{1},
					[]int{1, 2},
					[]int{1, 2, 2},
					[]int{1, 2, 2, 3},
					[]int{1, 2, 3},
					[]int{1, 3},
					[]int{2},
					[]int{2, 2},
					[]int{2, 2, 3},
					[]int{2, 3},
					[]int{3},
				},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, subsetsWithDup(p.nums), "输入:%v", p)
	}
}

func Benchmark_subsetsWithDup(b *testing.B) {
	A := []int{1, 2, 3, 6, 5, 4, 7, 8, 9, 5, 6, 3, 2, 4}
	for i := 1; i < b.N; i++ {
		subsetsWithDup(A)
	}
}
