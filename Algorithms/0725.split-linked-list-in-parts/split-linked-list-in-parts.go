package problem0725

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type ListNode = kit.ListNode

func splitListToParts(root *ListNode, k int) []*ListNode {
	size := length(root)
	remainder := size % k

	res := make([]*ListNode, k)

	i := 0
	for {
		res[i] = root
		i++
		if i == k {
			break
		}

		leng := size / k
		if remainder > 0 {
			leng++
			remainder--
		}

		for leng > 1 && root != nil {
			root = root.Next
			leng--
		}

		if root == nil {
			// 当 size < k 时，会出现这种情况
			break
		}

		// 斩断链条
		root.Next, root = nil, root.Next
	}

	return res
}

func length(root *ListNode) int {
	res := 0
	for root != nil {
		res++
		root = root.Next
	}
	return res
}
