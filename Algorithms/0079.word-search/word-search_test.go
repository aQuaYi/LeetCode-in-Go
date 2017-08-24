package Problem0079

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
	word string
}

// ans 是答案
type ans struct {
	one bool
}

func Test_Problem0079(t *testing.T) {
	ast := assert.New(t)
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	qs := []question{

		question{
			para{
				"ABCCED",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"SEE",
			},
			ans{
				true,
			},
		},

		question{
			para{
				"ABCB",
			},
			ans{
				false,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, exist(board, p.word), "输入:%v", p)
	}

	ast.False(exist([][]byte{}, ""))

	ast.False(exist([][]byte{[]byte{}}, ""))

	ast.False(exist([][]byte{[]byte("ABC")}, ""))
}
