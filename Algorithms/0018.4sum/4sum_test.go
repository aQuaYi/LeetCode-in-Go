package problem0018

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
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one [][]int
}

func Test_Problem0018(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans{[][]int{
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			}},
		},
		question{
			para{[]int{0, 0, 0, 0}, 0},
			ans{[][]int{
				[]int{0, 0, 0, 0},
			}},
		},
		question{
			para{[]int{-1, -1, -2, -2, 1, 1, 2, 2, 0, 0, 0, 0}, 0},
			ans{[][]int{
				[]int{-2, -2, 2, 2},
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-2, 0, 1, 1},
				[]int{-1, -1, 0, 2},
				[]int{-1, -1, 1, 1},
				[]int{-1, 0, 0, 1},
				[]int{0, 0, 0, 0},
			}},
		},
		question{
			para{[]int{0, -1, 0, 1, -2, -5, 3, 5, 0}, 6},
			ans{[][]int{
				[]int{-2, 0, 3, 5},
				[]int{0, 0, 1, 5},
			}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, fourSum(p.one, p.two), "输入:%v", p)
		printIntss(a.one, fourSum(p.one, p.two))
	}
}

func printIntss(e, a [][]int) {
	fmt.Println("expected \tactual")
	l := len(e)
	if l < len(a) {
		l = len(a)
	}
	for i := 0; i < l; i++ {
		if i < len(e) {
			fmt.Print(e[i], "\t")
		} else {
			fmt.Print("\t\t")
		}

		if i < len(a) {
			fmt.Print(a[i])
		}
		fmt.Println()
	}
}
