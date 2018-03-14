package problem0086

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
	x    int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_Problem0086(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 4, 3, 2, 5, 2},
				3,
			},
			ans{
				[]int{1, 2, 2, 4, 3, 5},
			},
		},

		question{
			para{
				[]int{1, 1},
				0,
			},
			ans{
				[]int{1, 1},
			},
		},

		question{
			para{
				[]int{1},
				3,
			},
			ans{
				[]int{1},
			},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(partition(s2l(p.head), p.x)), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}
	length := 10

	for head != nil {
		length--
		res = append(res, head.Val)
		head = head.Next
		if length < 0 {
			panic("链表长度超过10，可能出错")
		}
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
