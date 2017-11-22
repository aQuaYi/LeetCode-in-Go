package Problem0556

import "sort"

func nextGreaterElement(n int) int {
	nums := make([]int, 0, 10)

	lastTail := n % 10
	isAvaliable := false
	for n > 0 {
		tail := n % 10
		if tail < lastTail {
			isAvaliable = true
		}
		lastTail = tail

		nums = append(nums, tail)

		n /= 10
	}

	if !isAvaliable {
		return -1
	}

	reverse(nums)

	beg := exchange(nums)

	sort.Ints(nums[beg:])

	s := combine(nums)

	if s > 1<<31-1 {
		return -1
	}
	return s
}

func reverse(ss []int) {
	i, j := 0, len(ss)-1
	for i < j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
}

func exchange(ss []int) int {
	var i, j int
	maxStr := 10

	for i = len(ss) - 2; 0 <= i; i-- {
		n := ss[i]
		minGreater := maxStr
		minIndex := i
		for j = i + 1; j < len(ss); j++ {
			if n < ss[j] && ss[j] < minGreater {
				minGreater = ss[j]
				minIndex = j
			}
		}

		if i < minIndex {
			ss[i], ss[minIndex] = ss[minIndex], ss[i]
			return i + 1
		}
	}

	// 因为主函数前面的有 isAValiable 检查过了
	// 所以，总会出现交换的情况
	panic("NEVER BE HERE")
}

func combine(ss []int) int {
	str := 0
	for i := range ss {
		str = str*10 + ss[i]
	}
	return str
}
