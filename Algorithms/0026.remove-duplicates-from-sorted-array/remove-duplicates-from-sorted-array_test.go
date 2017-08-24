package Problem0026

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
	two []int
}

func Test_Problem0026(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 2, 2, 3, 3, 4, 5}},
			ans{5, []int{1, 2, 3, 4, 5}},
		},

		question{
			para{[]int{1}},
			ans{1, []int{1}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		l := removeDuplicates(p.one)
		ast.Equal(a.one, l, "输入:%v", p)
		ast.Equal(a.two, p.one[:l])
	}
}
