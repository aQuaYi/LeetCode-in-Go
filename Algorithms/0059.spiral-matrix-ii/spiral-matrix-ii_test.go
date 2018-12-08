package problem0059

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
	n int
}

// ans 是答案
type ans struct {
	one [][]int
}

func Test_Problem0059(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{0},
			ans{
				nil,
			},
		},

		question{
			para{3},
			ans{
				[][]int{
					[]int{1, 2, 3},
					[]int{8, 9, 4},
					[]int{7, 6, 5},
				},
			},
		},

		question{
			para{4},
			ans{
				[][]int{
					[]int{1, 2, 3, 4},
					[]int{12, 13, 14, 5},
					[]int{11, 16, 15, 6},
					[]int{10, 9, 8, 7},
				},
			},
		},

		question{
			para{5},
			ans{
				[][]int{
					[]int{1, 2, 3, 4, 5},
					[]int{16, 17, 18, 19, 6},
					[]int{15, 24, 25, 20, 7},
					[]int{14, 23, 22, 21, 8},
					[]int{13, 12, 11, 10, 9},
				},
			},
		},

		question{
			para{6},
			ans{
				[][]int{
					[]int{1, 2, 3, 4, 5, 6},
					[]int{20, 21, 22, 23, 24, 7},
					[]int{19, 32, 33, 34, 25, 8},
					[]int{18, 31, 36, 35, 26, 9},
					[]int{17, 30, 29, 28, 27, 10},
					[]int{16, 15, 14, 13, 12, 11},
				},
			},
		},

		question{
			para{7},
			ans{
				[][]int{
					[]int{1, 2, 3, 4, 5, 6, 7},
					[]int{24, 25, 26, 27, 28, 29, 8},
					[]int{23, 40, 41, 42, 43, 30, 9},
					[]int{22, 39, 48, 49, 44, 31, 10},
					[]int{21, 38, 47, 46, 45, 32, 11},
					[]int{20, 37, 36, 35, 34, 33, 12},
					[]int{19, 18, 17, 16, 15, 14, 13},
				},
			},
		},

		question{
			para{8},
			ans{
				[][]int{
					[]int{1, 2, 3, 4, 5, 6, 7, 8},
					[]int{28, 29, 30, 31, 32, 33, 34, 9},
					[]int{27, 48, 49, 50, 51, 52, 35, 10},
					[]int{26, 47, 60, 61, 62, 53, 36, 11},
					[]int{25, 46, 59, 64, 63, 54, 37, 12},
					[]int{24, 45, 58, 57, 56, 55, 38, 13},
					[]int{23, 44, 43, 42, 41, 40, 39, 14},
					[]int{22, 21, 20, 19, 18, 17, 16, 15},
				},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, generateMatrix(p.n), "输入:%v", p)
	}
}
