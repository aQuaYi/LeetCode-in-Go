package problem0119

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
	rowIndex int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0119(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				3,
			},
			ans{
				[]int{1, 3, 3, 1},
			},
		},
		
		question{
			para{
				0,
			},
			ans{
				[]int{1},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, getRow(p.rowIndex), "输入:%v", p)
	}
}
