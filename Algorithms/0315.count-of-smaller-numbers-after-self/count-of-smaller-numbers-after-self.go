package problem0315

type Node struct {
	left, right         *Node
	val, smaller, count int
}

func (root *Node) insert(val int) int {
	sum := 0

	for root.val != val {
		if root.val > val {
			if root.left == nil {
				root.left = &Node{
					val: val,
				}
			}

			root.smaller++
			root = root.left
		} else {
			if root.right == nil {
				root.right = &Node{
					val: val,
				}
			}

			sum += root.smaller + root.count
			root = root.right
		}
	}

	root.count++
	return sum + root.smaller
}

func countSmaller(nums []int) []int {
	size := len(nums)
	res := make([]int, size)

	if size <= 1 {
		return res
	}

	root := &Node{
		val: nums[size-1],
	}

	for i := size - 1; i >= 0; i-- {
		res[i] = root.insert(nums[i])
	}

	return res
}
