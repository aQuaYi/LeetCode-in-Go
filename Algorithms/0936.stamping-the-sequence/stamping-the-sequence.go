package problem0936

import "strings"

// ref: https://leetcode.com/problems/stamping-the-sequence/discuss/189254/Python-Greedy-and-DFS

func movesToStamp(stamp string, target string) []int {
	s, t := []byte(stamp), []byte(target)
	sSize, tSize := len(stamp), len(target)

	res := make([]int, 0, 1000)

	// isStamped return true if stamped since i
	isStamped := func(i int) bool {
		canStamp := false
		for j := 0; j < sSize; j++ {
			if t[i+j] == '?' {
				continue
			}
			if t[i+j] != s[j] {
				return false
			}
			canStamp = true
		}
		if canStamp {
			for k := i; k < i+sSize; k++ {
				t[k] = '?'
			}
			res = append(res, i)
		}
		return canStamp
	}

	maxIndex := tSize - sSize + 1

	for {
		isChanged := false
		for i := 0; i < maxIndex; i++ {
			isChanged = isChanged || isStamped(i)
		}
		if !isChanged {
			// no more place to stamp
			break
		}
	}

	reverse(res)

	if string(t) == strings.Repeat("?", tSize) {
		return res
	}
	return nil
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
