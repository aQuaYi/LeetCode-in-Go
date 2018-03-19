package problem0779

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}

	if K%2 == 1 {
		//      0
		//     / \
		//    0   1
		//   /\   /\
		//  0  1 1  0
		// 把变化按照 binary tree 的样子写下来，会发现
		// 当 K 为奇数时，其数值与其父节点 (N-1, (K+1)/2) 一样
		return kthGrammar(N-1, (K+1)/2)
	}

	// 当 K 为偶数时，其数值与其父节点 (N-1, K/2) 相反
	return opposite(kthGrammar(N-1, K/2))
}

// 当 a 是 0 时，返回 1
// 当 a 是 1 时，返回 0
func opposite(a int) int {
	return (a + 1) % 2
}
