package problem0834

// Input: N = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
// Output: [8,12,6,10,10,10]
// Explanation:
// Here is a diagram of the given tree:
//   0
//  / \
// 1   2
//    /|\
//   3 4 5
// We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
// equals 1 + 1 + 2 + 2 + 2 = 8.  Hence, answer[0] = 8, and so on.

// dist(0,5) 可以看做是线段 [0,5] 的长度。

func sumOfDistancesInTree(N int, edges [][]int) []int {
	tree := make([][]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}

	res := make([]int, N)
	count := make([]int, N)

	isSeen := make([]bool, N)
	deepFirstSearch(0, isSeen, count, res, tree)

	// 此时，count[i]=x 表示
	// 以 0 为 root 节点的大树中的
	// 以 i 为 root 的子树中，一共含有 x 个节点。

	// 此时，只有 res[0] 中的值是正确的值，因为 dfs 从 0 节点开始搜索

	isSeen = make([]bool, N)
	deepFirstMove(0, isSeen, count, res, tree)

	return res
}

func deepFirstSearch(root int, isSeen []bool, count, res []int, tree [][]int) {
	isSeen[root] = true
	for _, n := range tree[root] {
		if isSeen[n] {
			continue
		}
		deepFirstSearch(n, isSeen, count, res, tree)
		count[root] += count[n]
		res[root] += res[n] + count[n]
		// 上句的含义是
		// 所有线段 [root, x]，(x 为 n 为根节点的子树中的任意节点) 的长度之和，为
		// (res[n]+count[n]-1) + 1
		// 		^ 				 ^
		// 		|   			 |
		// 		|				新线段 [root,n]，长度为 1
		// 	 n 子树中，所有线段长度之和为 res[n]，
		// 	 原先有 count[n]-1 的线段，每根的长度
		// 	 增加 1。
	}
	count[root]++
}

// 从 root 开始
// 把 root 节点的子节点，设置成新的根节点
func deepFirstMove(root int, isSeen []bool, count, res []int, tree [][]int) {
	N := len(isSeen)
	isSeen[root] = true
	for _, n := range tree[root] {
		if isSeen[n] {
			continue
		}
		// 根节点由 root 变成 n 后
		res[n] = res[root] + (N - count[n]) - count[n]
		//			^			^				^
		// 			|			|				|
		// 			|			|			 n 子树中节点为端点的线段长度，全部 -1
		//			|		以飞 n 子树中节点为端点的线段长度，全部 +1
		// 		变化根节点前，全部线段的长度
		deepFirstMove(n, isSeen, count, res, tree)
	}
}
