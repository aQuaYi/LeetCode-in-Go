package problem0310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// 所有与 i 节点相连接的节点名称，保存在 linkNodes[i]
	linkNodes := make([][]int, n)
	for i := 0; i < n; i++ {
		linkNodes[i] = make([]int, 0, 5)
	}

	// count[i] 保存了与 i 节点链接的点的个数
	count := make([]int, n)
	for _, e := range edges {
		linkNodes[e[0]] = append(linkNodes[e[0]], e[1])
		linkNodes[e[1]] = append(linkNodes[e[1]], e[0])
		count[e[0]]++
		count[e[1]]++
	}

	// 收集所有的 叶子 节点
	var leafs = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if count[i] == 1 {
			leafs = append(leafs, i)
		}
	}

	var nd, leaf int

	// 反复裁剪掉 叶子 节点。
	// 当剩余节点数 <=2 时，就是答案
	for n > 2 {
		newLeafs := make([]int, 0, len(leafs))
		for _, leaf = range leafs {
			n--
			for _, nd = range linkNodes[leaf] {
				count[nd]--
				if count[nd] == 1 {
					newLeafs = append(newLeafs, nd)
				}
			}
		}
		leafs = newLeafs
	}

	var res = make([]int, 0, len(leafs))
	for _, v := range leafs {
		res = append(res, v)
	}

	return res
}
