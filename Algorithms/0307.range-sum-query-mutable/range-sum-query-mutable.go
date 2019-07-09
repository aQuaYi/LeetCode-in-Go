package problem0307

// https://leetcode.com/problems/range-sum-query-mutable/discuss/75724/17-ms-Java-solution-with-segment-tree
type SegmentTreeNode struct {
	start, end, sum int
	left, right     *SegmentTreeNode
}

func (node *SegmentTreeNode) update(i, val int) {
	if node.start == node.end {
		node.sum = val
		return
	}

	mid := node.start + (node.end-node.start)/2
	if mid >= i {
		node.left.update(i, val)
	} else {
		node.right.update(i, val)
	}

	node.sum = node.left.sum + node.right.sum
}

func (node *SegmentTreeNode) sumRange(i int, j int) int {
	if node.start == i && node.end == j {
		return node.sum
	}

	mid := node.start + (node.end-node.start)/2
	if mid >= j {
		return node.left.sumRange(i, j)
	} else if mid < i {
		return node.right.sumRange(i, j)
	} else {
		return node.left.sumRange(i, mid) + node.right.sumRange(mid+1, j)
	}
}

func buildTree(nums []int, start, end int) *SegmentTreeNode {
	if start > end {
		return nil
	}

	node := &SegmentTreeNode{
		start: start,
		end:   end,
	}

	if start == end {
		node.sum = nums[start]
	} else {
		mid := start + (end-start)/2
		node.left = buildTree(nums, start, mid)
		node.right = buildTree(nums, mid+1, end)
		node.sum = node.left.sum + node.right.sum
	}

	return node
}

type NumArray struct {
	root *SegmentTreeNode
}

func Constructor(nums []int) NumArray {
	return NumArray{
		buildTree(nums, 0, len(nums)-1),
	}
}

func (this *NumArray) Update(i int, val int) {
	this.root.update(i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.root.sumRange(i, j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
