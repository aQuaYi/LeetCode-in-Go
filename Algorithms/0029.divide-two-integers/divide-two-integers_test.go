package problem0029

import (
	"fmt"
	"math"
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
	one int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem0029(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{2, 3},
			ans{0},
		},

		question{
			para{2, 0},
			ans{math.MaxInt32},
		},

		question{
			para{1024, 3},
			ans{341},
		},

		question{
			para{1024, -3},
			ans{-341},
		},

		question{
			para{-1024, 3},
			ans{-341},
		},

		question{
			para{-1024, -3},
			ans{341},
		},

		question{
			para{1024, 1},
			ans{1024},
		},

		question{
			para{2147483647, 1},
			ans{2147483647},
		},

		question{
			para{2147483648, 1},
			ans{math.MaxInt32},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, divide(p.one, p.two), "输入:%v", p)
	}
}
func Test_d(t *testing.T) {
	ast := assert.New(t)

	res, remainder := d(1024, 3, 1)
	ast.Equal(1, remainder, "余数不对")
	ast.Equal(341, res, "结果不对")

	res, remainder = d(2147483647, 1, 1)
	ast.Equal(0, remainder, "余数不对")
	ast.Equal(2147483647, res, "结果不对")
}
