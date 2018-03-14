package problem0552

var m = 1000000007

func checkRecord(n int) int {
	// 为了后面的迭代，先分类讨论一下 n
	switch n {
	case 0:
		return 1
	case 1:
		return 3
	case 2: // 后面初始化 A 的时候，会有 A[2]
		return 8
	}

	// P[i] 表示，共有 i 个记录，且第 i 个记录为 P 的排列数
	P := make([]int, n)
	P[0] = 1
	// L[i] 表示，共有 i 个记录，且第 i 个记录为 L 的排列数
	L := make([]int, n)
	L[0], L[1] = 1, 3
	// A[i] 表示，共有 i 个记录，且第 i 个记录为 A 的排列数
	A := make([]int, n)
	A[0], A[1], A[2] = 1, 2, 4
	// 由以上 P，L，A 的定义可知
	// 所求的答案 T[n-1] == P[n-1] + L[n-1] + A[n-1]

	for i := 1; i < n; i++ {
		// 为了避免 P, L, A 中的数字累计过大
		// 对 i-1 上的数进行求余运算
		P[i-1] %= m
		L[i-1] %= m
		A[i-1] %= m

		// 当第 i 个为 'P' 时
		P[i] = P[i-1] + // i-1 可以为 P
			L[i-1] + // i-1 可以为 L
			A[i-1] // i-1 可以为 A

		if i > 1 {
			// 当第 i 个为 'L' 时
			L[i] = P[i-1] + // i-1 可以为 P
				P[i-2] + A[i-2] + // i-1 可以为 L，但是 i-2 必须为 P 或 A
				A[i-1] // i-1 可以为 A
		}

		if i > 2 {
			// 关于下面的表达式，我没有想到数学上的合理解释
			// 具体来源请参考
			// https://discuss.leetcode.com/topic/86696/share-my-o-n-c-dp-solution-with-thinking-process-and-explanation/3
			A[i] = A[i-1] + A[i-2] + A[i-3]
		}
	}

	return (P[n-1] + L[n-1] + A[n-1]) % m
}
