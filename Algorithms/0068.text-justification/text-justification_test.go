package problem0068

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
	words    []string
	maxWidth int
}

// ans 是答案
type ans struct {
	one []string
}

func Test_Problem0068(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]string{"This", "is", "an", "example", "of", "text", "justification."},
				16,
			},
			ans{
				[]string{"This    is    an", "example  of text", "justification.  "},
			},
		},

		question{
			para{
				[]string{"What", "must", "be", "shall", "be."},
				12,
			},
			ans{
				[]string{"What must be", "shall be.   "},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, fullJustify(p.words, p.maxWidth), "输入:%v", p)
	}
}
