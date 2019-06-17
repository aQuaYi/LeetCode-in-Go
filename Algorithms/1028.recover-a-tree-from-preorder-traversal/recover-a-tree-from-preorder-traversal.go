package problem1028

import (
	"strconv"
	"strings"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func recoverFromPreorder(S string) *TreeNode {
	m := make(map[int][]*TreeNode, 100)

	index := strings.IndexByte(S, '-')
	if index == -1 {
		val, _ := strconv.Atoi(S)
		return &TreeNode{
			Val: val,
		}
	}

	v := S[:index]
	S = S[index:]
	val, _ := strconv.Atoi(v)
	m[0] = append(m[0], &TreeNode{
		Val: val,
	})

	level, val := 0, 0
	for S != "" {
		level, val, S = next(S)
		node := &TreeNode{Val: val}
		m[level] = append(m[level], node)
		p := m[level-1][len(m[level-1])-1]
		if p.Left == nil {
			p.Left = node
		} else {
			p.Right = node
		}
	}

	return m[0][0]
}

func next(S string) (int, int, string) {
	level := 0
	for S[level] == '-' {
		level++
	}
	end := level
	for end < len(S) && S[end] != '-' {
		end++
	}
	value, _ := strconv.Atoi(S[level:end])
	return level, value, S[end:]
}
