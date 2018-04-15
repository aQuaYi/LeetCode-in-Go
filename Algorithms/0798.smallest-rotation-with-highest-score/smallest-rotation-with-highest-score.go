package problem0798

// NOTICE: a[i] 的范围是 [0, len(a))
// 所以，每次移动一格的话
// a[0] 中的元素移动到了末尾 score++，只有这个会带来 score 加分
// 同时，所有移动前 a[i]==i 的元素，会导致 score--，只有这个会带来 score 减分
// 所以，不用计算具体的 score 是多少
// 只要比较 score[k] = score[0] + delta[k] 中的 delta[k] 就好了

func bestRotation(a []int) int {
	size := len(a)
	delta := make([]int, size)
	for i := 0; i < size; i++ {
		delta[(i-a[i]+1+size)%size]--
	}
	// 此时，假设 delta[3] == -2
	// 表示，每次移动一位的话
	// 第 3 次移动前，有 2 个元素是处于 a[i]==i 的状态，
	// 所以，第 3 次移动后， score 需要减去 2 分

	maxIdx := 0
	for k := 1; k < size; k++ {
		// delta[k] = delta[k] + delta[k-1] + 1
		//    ^         ^          ^          ^
		//    |         |          |          |
		//    |         |          |         第 k 次移动后，a[0] 移到末尾带来的加分
		//    |         |        delta[k-1] = score[k-1] - score[0]
		//    |       第 k 次移动后，会带来的分数减少
		// delta[k] = score[k] - score[0]
		delta[k] += delta[k-1] + 1
		if delta[maxIdx] < delta[k] {
			maxIdx = k
		}
	}

	return maxIdx
}
