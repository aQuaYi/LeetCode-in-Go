package problem0024

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
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_Problem0024(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{[]int{}},
			ans{[]int{}},
		},

		question{
			para{[]int{1}},
			ans{[]int{1}},
		},

		question{
			para{[]int{1, 2, 3, 4}},
			ans{[]int{2, 1, 4, 3}},
		},

		question{
			para{[]int{1, 2, 3, 4, 5}},
			ans{[]int{2, 1, 4, 3, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(swapPairs(s2l(p.one))), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func Benchmark_swapPairs(b *testing.B) {
	head := s2l([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for i := 1; i < b.N; i++ {
		swapPairs(head)
	}
}
