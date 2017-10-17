package Problem0315

type treeNode struct {
	num, idx, count int
	left, right     *treeNode
}

func countSmaller(nums []int) []int {
	root := buildTree(nums)

	res := make([]int, len(nums))

	var search func(*treeNode)
	search = func(root *treeNode) {
		if root == nil {
			return
		}
		res[root.idx] = root.count
		search(root.left)
		search(root.right)
	}

	search(root)

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
