package problem0057

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
	intervals   []Interval
	newInterval Interval
}

// ans 是答案
type ans struct {
	one []Interval
}

func Test_Problem0057(t *testing.T) {
	ast := assert.New(t)
	qs := []question{

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{7, 8},
			},
			ans{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{4, 5},
			},
			ans{[]Interval{
				Interval{1, 3},
				Interval{4, 5},
				Interval{6, 9},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{2, 5},
			},
			ans{[]Interval{
				Interval{1, 5},
				Interval{6, 9},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{-1, 0},
			},
			ans{[]Interval{
				Interval{-1, 0},
				Interval{1, 3},
				Interval{6, 9},
			},
			},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{11, 20},
			},
			ans{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
				Interval{11, 20},
			},
			},
		},

		question{
			para{[]Interval{},
				Interval{2, 5},
			},
			ans{[]Interval{
				Interval{2, 5},
			},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, insert(p.intervals, p.newInterval), "输入:%v", p)
	}
}

func Test_min(t *testing.T) {
	ast := assert.New(t)

	actual := min(1, 0)
	expected := 0
	ast.Equal(expected, actual)

}

func Test_max(t *testing.T) {
	ast := assert.New(t)

	actual := max(1, 0)
	expected := 1
	ast.Equal(expected, actual)

}
