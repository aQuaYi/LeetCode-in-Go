package problem0033

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
	one int
}

func Test_Problem0033(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{2}, 2},
			ans{0},
		},

		question{
			para{[]int{8, 9, 2, 3, 4}, 9},
			ans{1},
		},

		question{
			para{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans{4},
		},

		question{
			para{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 1},
			ans{3},
		},

		question{
			para{[]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2}, 5},
			ans{1},
		},

		question{
			para{[]int{1, 3}, 0},
			ans{-1},
		},

		question{
			para{[]int{}, 5},
			ans{-1},
		},

		question{
			para{[]int{2}, 5},
			ans{-1},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, search(p.one, p.two), "输入:%v", p)
	}
}

func Test_indexOfMin_evenLength(t *testing.T) {
	ast := assert.New(t)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsSize := 10
	for i := 0; i < 10; i++ {
		nums := array[i : i+numsSize]
		indexOfZeroInNums := 9 - i
		actual := indexOfMin(nums)
		ast.Equal(indexOfZeroInNums, actual)
	}
}

func Test_indexOfMin_oddLength(t *testing.T) {
	ast := assert.New(t)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsSize := 9
	for i := 1; i < 10; i++ {
		nums := array[i : i+numsSize]
		indexOfZeroInNums := 9 - i
		actual := indexOfMin(nums)
		ast.Equal(indexOfZeroInNums, actual)
	}
}
