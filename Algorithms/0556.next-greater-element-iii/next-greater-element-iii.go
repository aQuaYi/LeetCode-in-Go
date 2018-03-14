package problem0556

import "sort"

func nextGreaterElement(n int) int {
	nums := make([]int, 0, 10)

	lastTail := n % 10
	// 用于标记 n 已经为其能表现的最大值
	// 例如, n == 4321，就不可能变得更大了
	isMax := true
	for n > 0 {
		tail := n % 10
		if tail < lastTail {
			// 较高位上存在较小的值
			// n 还可以变大
			isMax = false
		}
		lastTail = tail
		nums = append(nums, tail)
		n /= 10
	}

	// 由于 n 不可能再变大了，所以，提前结束
	if isMax {
		return -1
	}

	// nums 中 digit 的存入顺序与实际顺序相反
	// 需要逆转一下 nums
	reverse(nums)

	// 按照题意，交换 nums 中，两个数的位子
	beg := exchange(nums)

	// 重新排列 nums 尾部的数字，使得 nums 更小
	sort.Ints(nums[beg:])

	res := combine(nums)

	if res > 1<<31-1 {
		return -1
	}
	return res
}

func reverse(ss []int) {
	i, j := 0, len(ss)-1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}

// 找到最大的 i 使得
// s[j] 是 s[i+1:] 中大于 s[i] 的最小值
// s[i], s[j] 互换后
// 返回 i+1
func exchange(a []int) int {
	var i, j int

	for i = len(a) - 2; 0 <= i; i-- {
		n := a[i]
		min := 10
		index := i
		for j = i + 1; j < len(a); j++ {
			if n < a[j] && a[j] < min {
				min = a[j]
				index = j
			}
		}

		if i < index {
			a[i], a[index] = a[index], a[i]
			break
		}
	}

	return i + 1
}

// 把 a 整合成一个 int
func combine(a []int) int {
	num := 0
	for i := range a {
		num = num*10 + a[i]
	}
	return num
}
