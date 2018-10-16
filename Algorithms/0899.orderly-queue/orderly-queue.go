package problem0899

import (
	"sort"
)

func orderlyQueue(s string, k int) string {
	if k == 1 {
		return minRotated(s)
	}
	// 当 k >=2 时
	// 前两个位置上的字母，总是把较小的字母移到末尾
	// 经过 len(s)-1 次比较后，s 中最大的字母，一定在前两个位置中，
	// 此时，把 s 中最大的字母移到末尾。
	// 以此类推，可以形成一个冒泡排序。
	// 所以，当 k >= 2 时，可以直接对 s 进行排序
	// 当 k = 1 时，无法进行比较
	// 就只能通过回转 s 来查看最小的字符串了
	return sorted(s)
}

func minRotated(s string) string {
	min := s
	bytes := []byte(s)
	bytes = append(bytes, bytes...)
	size := len(s)
	for i := 1; i < size; i++ {
		rs := string(bytes[i : i+size])
		if min > rs {
			min = rs
		}
	}
	return min
}

func sorted(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i int, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
