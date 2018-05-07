package problem0817

import (
	"sort"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * 	return 0
}
*/

// ListNode is from kit
type ListNode = kit.ListNode

func numComponents(head *ListNode, G []int) int {
	// 收集 list 中各个数的索引号
	// 方便以后判断连通性
	indexs := make(map[int]int, len(G)*2)
	i := 0
	for head != nil {
		indexs[head.Val] = i
		i++
		head = head.Next
	}

	// 初步统计连通性
	// temp[i][0] == 3 表示，第 i 个联通的链条 起点 是 3
	// temp[i][1] == 5 表示，第 i 个联通的链条 终点 是 5
	temp := make([][2]int, 0, len(G))
	for _, num := range G {
		idx := indexs[num]
		isConnected := false
		for i := range temp {
			// 如果 idx 与 temp[i] 的 起点 相连
			if idx+1 == temp[i][0] {
				temp[i][0] = idx
				isConnected = true
				break
			}
			// 如果 idx 与 temp[i] 的 终点 相连
			if temp[i][1] == idx-1 {
				temp[i][1] = idx
				isConnected = true
				break
			}
		}
		// 如果 idx 不与 temp 中的任何链条相连
		if !isConnected {
			temp = append(temp, [2]int{idx, idx})
		}
	}

	// 按照链条的起点排序
	sort.Slice(temp, func(i int, j int) bool {
		return temp[i][0] < temp[j][0]
	})

	t := temp[0]
	res := 1
	for i := 1; i < len(temp); i++ {
		// 如果前后两个链条不相连
		if t[1]+1 != temp[i][0] {
			res++
		}
		t = temp[i]
	}

	return res
}
