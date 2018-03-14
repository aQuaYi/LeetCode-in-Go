package problem0036

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
	one [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one bool
}

func Test_Problem0036(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]byte{
				[]byte(".87654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			ans{true},
		},

		question{
			para{[][]byte{
				[]byte(".87654321"),
				[]byte("2.3......"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			ans{false},
		},

		question{
			para{[][]byte{
				[]byte(".88654321"),
				[]byte("2........"),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			ans{false},
		},

		question{
			para{[][]byte{
				[]byte(".87654321"),
				[]byte("2....4..."),
				[]byte("3........"),
				[]byte("4........"),
				[]byte("5........"),
				[]byte("6........"),
				[]byte("7........"),
				[]byte("8........"),
				[]byte("9........"),
			}},
			ans{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isValidSudoku(p.one), "输入:%v", p)
	}
}
