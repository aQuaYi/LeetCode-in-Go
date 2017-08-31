package Problem0084

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
	heights []int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0084(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{2, 1, 5, 6, 2, 3},
			},
			ans{
				10,
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, largestRectangleArea(p.heights), "输入:%v", p)
	}
}
