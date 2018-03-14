package problem0063

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
	obstacleGrid [][]int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0063(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[][]int{
					[]int{1},
				},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[][]int{
					[]int{0, 0, 0},
					[]int{0, 1, 0},
					[]int{0, 0, 0},
				},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[][]int{
					[]int{0, 1, 0},
					[]int{1, 1, 0},
					[]int{0, 0, 0},
				},
			},
			ans{
				0,
			},
		},

		question{
			para{
				[][]int{},
			},
			ans{0},
		},

		question{
			para{
				[][]int{
					[]int{},
					[]int{},
				},
			},
			ans{0},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, uniquePathsWithObstacles(p.obstacleGrid), "输入:%v", p)
	}
}
