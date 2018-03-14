package problem0025

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
	one []int
}

func Test_Problem0025(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			ans{[]int{3, 2, 1, 4, 5}},
		},

		question{
			para{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			ans{[]int{1, 2, 3, 4, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(reverseKGroup(s2l(p.one), p.two)), "输入:%v", p)
	}
}
func Test_needReverse(t *testing.T) {
	head := s2l([]int{1, 2, 3, 4, 5, 6})
	begin, ok := needReverse(head, 4)
	assert.True(t, ok, "长度足够的链却提示不能逆转")
	assert.Equal(t, []int{1, 2, 3, 4}, l2s(head), "前链不对, 没有被斩断")
	assert.Equal(t, []int{5, 6}, l2s(begin), "后链不对")
}
func Test_reverse(t *testing.T) {
	first, last := reverse(s2l([]int{1, 2, 3}))
	assert.Equal(t, []int{3, 2, 1}, l2s(first), "无法逆转链")
	assert.Equal(t, 1, last.Val, "链尾部的值不对")

	Nil, _ := reverse(s2l([]int{}))
	assert.Nil(t, Nil, "无法逆转空链")
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
