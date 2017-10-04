package Problem0220

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	size := len(nums)
	if k == 0 || t < 0 || size <= 1 {
		return false
	}
	// t 可以为 0
	// t 需要当除数
	// t++ 后，避免了除0错误。
	// 但 <=t 变成了 <t
	t++
	// 注意，此时已经 t++ 了
	// |a-b|<t
	// -t < a - b < t
	// -1 < a/t - b/t < 1
	// nMap[a/t] = a
	nMap := make(map[int]int, size)
	var ni, m int
	for i := 0; i < size; i++ {
		ni = nums[i]
		m = ni / t
		if ni < 0 {
			// 3 / 4 * 4 == 0, 0 是 3 的左边的 4 的倍数
			// -1/ 4 * 4 == 0, 0 是 -1的右边的 4 的倍数
			// 所以，当除以正数 t 以后，被除数由于正负号的不同，他们的变动方向是不一样的。
			// 当 ni < 0 时，m--，可以统一除法后的变动方向。
			// 可以取消段程序，看看单元测试中，下面这段的结果
			// {
			// 	[]int{-1, 3},
			// 	1,
			// 	3,
			// 	false,
			// },
			m--
		}
		if _, ok := nMap[m]; ok {
			return true
		} else if n, ok := nMap[m-1]; ok && abs(ni, n) < t {
			return true
		} else if n, ok := nMap[m+1]; ok && abs(ni, n) < t {
			return true
		}

		nMap[m] = ni
		if i >= k {
			delete(nMap, nums[i-k]/t)
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
