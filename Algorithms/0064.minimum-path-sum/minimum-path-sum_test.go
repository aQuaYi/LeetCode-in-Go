package problem0064

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
	grid [][]int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0064(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			}},
			ans{21},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, minPathSum(p.grid), "输入:%v", p)
	}
}
func Test_min(t *testing.T) {
	ast := assert.New(t)

	actual := min(1, 0)
	expected := 0
	ast.Equal(expected, actual)

}
