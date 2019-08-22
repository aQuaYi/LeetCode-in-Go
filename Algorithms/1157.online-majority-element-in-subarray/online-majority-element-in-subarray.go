package problem1157

import (
	"math/rand"
	"sort"
)

// ref: https://leetcode.com/problems/online-majority-element-in-subarray/discuss/355848/Python-Binary-Search-%2B-Find-the-Majority-Element

// MajorityChecker is ..
type MajorityChecker struct {
	arr []int
	a2i map[int][]int
}

// Constructor is ...
func Constructor(arr []int) MajorityChecker {
	a2i := make(map[int][]int, 64)
	for i, a := range arr {
		a2i[a] = append(a2i[a], i)
	}
	return MajorityChecker{
		arr: arr,
		a2i: a2i,
	}
}

// Query is ...
func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	for i := 0; i < 20; i++ {
		index := rand.Intn(right-left+1) + left
		a := mc.arr[index]
		a2i := mc.a2i[a]
		n := len(a2i)
		l := sort.Search(n, func(i int) bool { return a2i[i] >= left })
		r := sort.Search(n, func(i int) bool { return a2i[i] > right })
		if r-l >= threshold {
			return a
		}
	}
	return -1
}
