package problem1028

import (
	"strconv"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func recoverFromPreorder(S string) *TreeNode {
	next := func() (int, *TreeNode) {
		level := 0
		for S[level] == '-' {
			level++
		}
		end := level
		for end < len(S) && S[end] != '-' {
			end++
		}
		val, _ := strconv.Atoi(S[level:end])
		S = S[end:]
		return level, &TreeNode{Val: val}
	}

	stack, top := make([]*TreeNode, 1000), -1

	for len(S) > 0 {
		level, node := next()
		for top >= level { // top is the level of stack[top]
			top--
		}
		if top >= 0 {
			if stack[top].Left == nil {
				stack[top].Left = node
			} else {
				stack[top].Right = node
			}
		}
		top++
		stack[top] = node
	}

	return stack[0]
}
