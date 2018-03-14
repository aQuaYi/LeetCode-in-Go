package problem0088

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
	m   int
	two []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0088(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{0}, 0, []int{1}, 1},
			ans{[]int{1}},
		},

		question{
			para{[]int{1, 0}, 1, []int{2}, 1},
			ans{[]int{1, 2}},
		},

		question{
			para{[]int{}, 0, []int{1}, 1},
			ans{[]int{}},
		},

		question{
			para{[]int{1, 3, 5, 7}, 4, []int{2, 4}, 2},
			ans{[]int{1, 2, 3, 4}},
		},

		question{
			para{[]int{1, 3, 5, 7}, 4, []int{2, 2}, 2},
			ans{[]int{1, 2, 2, 3}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		merge(p.one, p.m, p.two, p.n)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
