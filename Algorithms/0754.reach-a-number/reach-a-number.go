package problem0754

import (
	"math"
)

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	n := int(math.Ceil((math.Sqrt(8*float64(target)+1) - 1) / 2))

	sum := n * (n + 1) / 2

	if sum == target {
		return n
	}

	x := sum - target

	// sum = 1 + 2 + 3 + 4 + 5 + ... + (n-1) + n
	// 可知 x<n,
	// 当 x 是偶数的时候，在 [1, n) 中必然存在一个 x/2，使得
	// target = 1 + 2 + ... + (-x/2) + ... + n
	if x&1 == 0 {
		return n
	}

	// 若 x 是奇数，且 (n+1) 是奇数，
	// 设 sum' = sum + (n+1)
	// x' = sum' - target = (n+1) + x 是偶数
	if (n+1)&1 == 1 {
		return n + 1
	}

	// 若 x 是奇数，且 (n+1) 是偶数，
	// 设 sum" = sum + (n+1) + (n+2)
	// x" = sum" - target = (n+2) + (n+1) + x 是偶数
	return n + 2

	// 总之，想要步骤最小，就要 [1,n] 中的负数个数最少
	// 但是负数，只能带来 △sum 为偶数
	// 需要调整 sum 的奇偶性
}
