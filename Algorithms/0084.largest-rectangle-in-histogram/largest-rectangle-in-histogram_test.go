package problem0084

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
				[]int{1, 100, 101, 1},
			},
			ans{
				200,
			},
		},

		question{
			para{
				[]int{2, 2, 2, 2, 2},
			},
			ans{
				10,
			},
		},

		question{
			para{
				[]int{6, 5, 4, 3, 2, 1},
			},
			ans{
				12,
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5, 6},
			},
			ans{
				12,
			},
		},

		question{
			para{
				[]int{2, 0, 1, 0, 1, 0},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[]int{2, 1, 5, 6, 2, 3},
			},
			ans{
				10,
			},
		},

		question{
			para{
				[]int{4, 2, 0, 3, 2, 4, 3, 4},
			},
			ans{
				10,
			},
		},

		question{
			para{
				[]int{4, 3, 4, 2, 3, 0, 2, 4},
			},
			ans{
				10,
			},
		},

		question{
			para{
				[]int{},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{1},
			},
			ans{
				1,
			},
		},

		question{
			para{
				[]int{0, 0},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{9, 0},
			},
			ans{
				9,
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
