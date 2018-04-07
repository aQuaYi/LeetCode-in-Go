package problem0786

import (
	"sort"
)

// 二分法查找
func kthSmallestPrimeFraction(A []int, K int) []int {
	// 确保 A 处于升序
	sort.Ints(A)

	lo, hi := 0.0, 1.0
	for {
		mid := (lo + hi) / 2

		p, q, count := countUnder(mid, A)

		switch {
		case count < K:
			lo = mid
		case count > K:
			hi = mid
		default:
			return []int{p, q}
		}
	}
}

// 统计 A 中所有能组成的分数中，数值 <= max 的数量
// 其中 p/q 是这些分数中的最大值
func countUnder(max float64, A []int) (p, q, count int) {
	n := len(A)
	// 把 q 赋值为 1 是为了能够更新 p，q 的值
	p, q, count = 0, 1, 0

	// A[i] 和 p 表示分子
	// A[j] 和 q 表示分母

	for i := 0; i < n-1; i++ {
		// 当分母为 A[i] 保存不变时
		// 利用二分法找到 j
		// 使得 A[j:] 中的所有元素做分母时，
		// 都满足 A[i]/A[j] <= mid
		// 注意，A 是升序

		lo, hi := i, n
		// 此时，A[hi:] 就是要找的 A[j:] 的子集
		// 利用二分法不断缩小 [lo,hi]，直到 lo == hi 时，
		// A[hi:] 就是 A[j:]

		// lo < hi, [lo:hi] 不为空，还要继续查找
		for lo < hi {
			mid := (hi + lo) / 2
			if float64(A[i]) <= max*float64(A[mid]) {
				// 当 A[i]/A[mid] <= max 时
				// A[mid:] 中的元素都 可以 满足要求
				// 所以 hi 向下移动
				hi = mid
			} else {
				// 当 max < A[i]/A[mid] 时
				// A[:mid+1] 中的元素都 不能 满足要求
				// 所以 lo 向上移动
				lo = mid + 1
			}
		}
		j := hi

		// n - j 是 A[j:] 中元素的个数
		count += n - j

		// n - j > 0 说明 A[j:] 中的存在满足要求的元素
		// 这些元素中作为分母，与 A[i] 组成分数的最大值，
		// 就是 A[i]/A[j]
		// 如果 p/q < A[i]/A[j]，就更新 p 和 q
		if n-j > 0 && p*A[j] < q*A[i] {
			p, q = A[i], A[j]
		}
	}

	return p, q, count
}
