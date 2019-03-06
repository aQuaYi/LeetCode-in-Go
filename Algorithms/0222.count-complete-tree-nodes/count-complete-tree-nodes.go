package problem0222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	traverse(root, &count)
	return count
}

func traverse(node *TreeNode, count *int) {
	if node == nil {
		return
	}

	*count++

	traverse(node.Left, count)
	traverse(node.Right, count)
}
