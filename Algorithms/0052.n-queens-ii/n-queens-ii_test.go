package problem0052

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
	n int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0052(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{1},
			ans{1},
		},

		question{
			para{0},
			ans{0},
		},

		question{
			para{8},
			ans{92},
		},

		question{
			para{12},
			ans{14200},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, totalNQueens(p.n), "输入:%v", p)
	}
}
