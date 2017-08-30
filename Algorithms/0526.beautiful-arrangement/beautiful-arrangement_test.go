package Problem0526

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
	N int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0526(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				2,
			},
			ans{
				2,
			},
		},

		question{
			para{
				3,
			},
			ans{
				3,
			},
		},

		question{
			para{
				15,
			},
			ans{
				24679,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, countArrangement(p.N), "输入:%v", p)
	}
}
