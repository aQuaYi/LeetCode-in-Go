package problem0061

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
	head []int
	k    int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0061(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{
				[]int{},
				1,
			},
			ans{
				[]int{},
			},
		},

		question{
			para{
				[]int{},
				0,
			},
			ans{
				[]int{},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{
				[]int{4, 5, 1, 2, 3},
			},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				99,
			},
			ans{
				[]int{2, 3, 4, 5, 1},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(rotateRight(s2l(p.head), p.k)), "输入:%v", p)
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
