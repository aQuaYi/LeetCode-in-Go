package problem0037

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
	one [][]byte
}

func Test_Problem0036(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]byte{
				[]byte("..9748..."),
				[]byte("7........"),
				[]byte(".2.1.9..."),
				[]byte("..7...24."),
				[]byte(".64.1.59."),
				[]byte(".98...3.."),
				[]byte("...8.3.2."),
				[]byte("........6"),
				[]byte("...2759.."),
			}},
			ans{[][]byte{
				[]byte("519748632"),
				[]byte("783652419"),
				[]byte("426139875"),
				[]byte("357986241"),
				[]byte("264317598"),
				[]byte("198524367"),
				[]byte("975863124"),
				[]byte("832491756"),
				[]byte("641275983"),
			}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		solveSudoku(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
