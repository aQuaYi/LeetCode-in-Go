package problem0050

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
	x float64
	n int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one float64
}

func Test_Problem0050(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{0, 1},
			ans{0},
		},

		question{
			para{3, 2},
			ans{9},
		},

		question{
			para{0.5, -3},
			ans{8},
		},

		question{
			para{0.00001, 2147483647},
			ans{0},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, myPow(p.x, p.n), "输入:%v", p)
	}
}
