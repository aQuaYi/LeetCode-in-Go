package problem0058

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
	s string
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0058(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				"haha",
			},
			ans{
				4,
			},
		},

		question{
			para{
				"haha ",
			},
			ans{
				4,
			},
		},
		
		question{
			para{
				"",
			},
			ans{
				0,
			},
		},

		question{
			para{
				"Hello World",
			},
			ans{
				5,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, lengthOfLastWord(p.s), "输入:%v", p)
	}
}
