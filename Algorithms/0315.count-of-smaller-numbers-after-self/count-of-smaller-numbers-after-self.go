package Problem0315

type treeNode struct {
	num, idx, count int
	left, right     *treeNode
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				res[i]++
			}
		}
	}
	return res
}

func buildTree(nums []int) *treeNode {
	var root *treeNode
	for i, n := range nums {
		node := treeNode{
			num: n,
			idx: i,
		}
		if root == nil {
			root = &node
		} else {
			addNode(root, &node)
		}
	}

	// TODO: 遍历树，获取 idx 和 count

	return root
}

func addNode(root *treeNode, node *treeNode) {
	if node.num < root.num {
		root.count++
		countUp(root.right)
		if root.left == nil {
			root.left = node
		} else {
			addNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			addNode(root.right, node)
		}
	}
}

func countUp(root *treeNode) {
	if root == nil {
		return
	}
	root.count++
	countUp(root.left)
	countUp(root.right)
}
