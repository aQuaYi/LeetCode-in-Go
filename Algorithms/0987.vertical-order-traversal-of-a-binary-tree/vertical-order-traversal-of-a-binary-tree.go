package problem0987

import (
	"sort"

	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// TreeNode is pre-defined...
type TreeNode = kit.TreeNode

func verticalTraversal(root *TreeNode) [][]int {
	ds := make([]data, 0, 1024)
	dfs(root, 0, 0, &ds)

	sort.Slice(ds, func(i int, j int) bool {
		if ds[i].x == ds[j].x {
			if ds[i].y == ds[j].y {
				return ds[i].value < ds[j].value
			}
			return ds[i].y > ds[j].y
		}
		return ds[i].x < ds[j].x
	})

	res := make([][]int, 0, len(ds))

	x := ds[0].x
	values := make([]int, 0, 10)
	for _, d := range ds {
		if x == d.x {
			values = append(values, d.value)
			continue
		}
		res = append(res, values)
		values = append(make([]int, 0, 10), d.value)
		x = d.x
	}
	res = append(res, values)

	return res
}

type data struct {
	value, x, y int
}

func dfs(node *TreeNode, x, y int, dataSlice *[]data) {
	if node == nil {
		return
	}
	dfs(node.Left, x-1, y-1, dataSlice)
	*dataSlice = append(*dataSlice, data{
		value: node.Val,
		x:     x,
		y:     y,
	})
	dfs(node.Right, x+1, y-1, dataSlice)
}
