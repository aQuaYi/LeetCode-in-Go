package problem0495

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
	timeSeries []int
	duration   int
}

// ans 是答案
type ans struct {
	one int
}

func Test_Problem0495(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{},
				2,
			},
			ans{
				0,
			},
		},

		question{
			para{
				[]int{1, 4},
				2,
			},
			ans{
				4,
			},
		},

		question{
			para{
				[]int{1, 2},
				2,
			},
			ans{
				3,
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, findPoisonedDuration(p.timeSeries, p.duration), "输入:%v", p)
	}
}
