package problem0483

import (
	"math"
	"strconv"
)

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseUint(n, 10, 64)
	// num = k^m + k^(m-1) + ... + k + 1
	// 想要找到最小的 k
	// 可知 k 变小的时候，m 会变大
	// k 最小可以是 2，即是 二进制
	// k == 2 时，m == mMax
	mMax := int(math.Log2(float64(num)))

	// 从 mMax 开始往下检查，对应的 k 能否满足题意
	for m := mMax; m >= 1; m-- {
		// k^m < num = k^m + k^(m-1) + ... + k + 1
		// 展开 (k+1)^m 后，可知
		// (k+1)^m > num = k^m + k^(m-1) + ... + k + 1
		// 综上所示
		// k^m < num < (k+1)^m，即
		// k < num^(1/m) < k+1，即
		// k == int(num^(1/m))
		k := uint64(math.Pow(float64(num), 1.0/float64(m)))
		// 这样就确定了 k 的取值，只需要验证 k 是否满足题意即可
		if isFound(num, k, m) {
			return strconv.FormatUint(k, 10)
		}
	}

	return strconv.FormatUint(num-1, 10)
}

// 返回 num == k^m + k^(m-1) + ... + k + 1
func isFound(num, k uint64, m int) bool {
	sum := uint64(1)
	a := uint64(1)

	for i := 1; i <= m; i++ {
		a *= k
		sum += a
	}

	return sum == num
}
