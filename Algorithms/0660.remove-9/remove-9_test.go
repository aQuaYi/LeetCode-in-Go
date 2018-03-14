package problem0660

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
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0660(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{45928},
			ans{70001},
		},

		question{
			para{100000000},
			ans{228145181},
		},

		question{
			para{8},
			ans{8},
		},

		question{
			para{9},
			ans{10},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, newInteger(p.one), "输入:%v", p)
	}
}
