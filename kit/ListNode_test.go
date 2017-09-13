package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_l2s(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{}, List2Slice(nil), "输入nil，没有返回[]int{}")

	one2three := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	ast.Equal([]int{1, 2, 3}, List2Slice(one2three), "没有成功地转换成[]int")

	limit := 100
	overLimitList := Slice2List(make([]int, limit+1))
	ast.Panics(func() { List2Slice(overLimitList) }, "转换深度超过 %d 限制的链条，没有 panic", limit)
}

func Test_s2l(t *testing.T) {
	ast := assert.New(t)
	ast.Nil(Slice2List([]int{}), "输入[]int{}，没有返回nil")

	ln := Slice2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	i := 1
	for ln != nil {
		ast.Equal(i, ln.Val, "对应的值不对")

		ln = ln.Next
		i++
	}
}
