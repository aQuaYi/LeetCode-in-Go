package problem0968

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

type status int

// every node just has one status of three
const (
	hasACamera status = iota
	isMonitoredByChild
	isMonitoredByParent
)

func minCameraCover(root *TreeNode) int {
	res := 0
	s := check(root, &res)
	if s == isMonitoredByParent {
		// root has no parent
		// so need itself has a camera
		res++
	}
	return res
}

func check(root *TreeNode, res *int) status {
	if root == nil {
		// nil don't need be monitored.
		// so we think it's monitored by his child
		return isMonitoredByChild
	}
	l, r := check(root.Left, res), check(root.Right, res)
	switch {
	case l == isMonitoredByParent || r == isMonitoredByParent:
		*res++
		return hasACamera
	case l == hasACamera || r == hasACamera:
		return isMonitoredByChild
	default:
		return isMonitoredByParent
	}
}
