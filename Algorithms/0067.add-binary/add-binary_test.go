package problem0067

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
	a string
	b string
}

// ans 是答案
type ans struct {
	one string
}

func Test_Problem0067(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"11",
				"1",
			},
			ans{
				"100",
			},
		},

		question{
			para{
				"001",
				"000",
			},
			ans{
				"1",
			},
		},

		question{
			para{
				"",
				"",
			},
			ans{
				"0",
			},
		},

		question{
			para{
				"",
				"1",
			},
			ans{
				"1",
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, addBinary(p.a, p.b), "输入:%v", p)
	}
}
