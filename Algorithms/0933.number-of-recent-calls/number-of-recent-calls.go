package problem0933

import "sort"

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

// RecentCounter ...
type RecentCounter struct {
	times []int
}

// Constructor ...
func Constructor() RecentCounter {
	return RecentCounter{
		times: make([]int, 0, 10000),
	}
}

// Ping use binary search
func (rc *RecentCounter) Ping(t int) int {
	rc.times = append(rc.times, t)
	return len(rc.times) - sort.SearchInts(rc.times, t-3000)
}
