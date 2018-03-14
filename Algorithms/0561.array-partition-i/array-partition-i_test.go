package problem0561

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

func Test_Problem0561(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{4, 3, 2, 1}},
			ans{4},
		},

		question{
			para{[]int{6, 7, 8, 5, 4, 3, 2, 1}},
			ans{16},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, arrayPairSum(p.one), "输入:%v", p)
	}
}
