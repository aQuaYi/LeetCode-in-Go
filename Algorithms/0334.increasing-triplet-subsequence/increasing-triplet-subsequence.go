package problem0334

func increasingTriplet(a []int) bool {
	max := 1<<63 - 1
	ai, aj := max, max

	for _, ak := range a {
		if ak <= ai {
			ai = ak
		} else if ak <= aj {
			aj = ak
		} else {
			return true
		}
	}

	return false
}

// TODO: 复习此题

// 按照题目要求，ai < aj < ak, 对应的索引值 i < j < k
// 按照程序的运行逻辑可知，ai < aj < ak，但对应的索引值，却有可能出现 i > j 的情况
// 例如，下面的 test case 返回 true 时
// [1,2,-10,3,-6]
//    ↑  ↑  ↑
//   aj ai ak
// 出现这样的情况，不代表程序出错了。因为
// 虽然，ai 从 1 变成了 -10 前，递增数的个数为 2，变化后，并没有让递增数的个数降低。
// 反过来说，保持递增数个数不变的前提下，ai 变小后，可以增大 aj 和 ak 的取值范围，避免遗漏某些情况。
// 观察下面 test case 的运行过程，可以清楚这一点。
// {[]int{1, 2, -10, -8, -6}, true},
