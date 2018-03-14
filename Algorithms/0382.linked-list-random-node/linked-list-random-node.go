package problem0382

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"math/rand"
)

type ListNode = kit.ListNode

// Solution 是需要设计的数据结构
// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }
type Solution struct {
	head *ListNode
	n    int
}

// Constructor 构建 Solution
// head is The linked list's head.
// Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	count := 1
	temp := head
	for temp.Next != nil {
		count++
		temp = temp.Next
	}
	return Solution{
		head: head,
		n:    count,
	}
}

// GetRandom returns a random node's value. */
func (s *Solution) GetRandom() int {
	count := rand.Intn(s.n)
	temp := s.head
	for count > 0 {
		temp = temp.Next
		count--
	}
	return temp.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
