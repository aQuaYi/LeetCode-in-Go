package Problem0205

import (
	"sort"
)

func isIsomorphic(s1, s2 string) bool {
	ints1 := makeInts(s1)
	ints2 := makeInts(s2)
	for i := 255; i >= 0; i-- {
		if ints1[i] == 0 && ints2[i] == 0 {
			break
		}
		if ints1[i] != ints2[i] {
			return false
		}
	}
	return true
}

func makeInts(s string) []int {
	ints := make([]int, 256)

	for i := 0; i < len(s); i++ {
		ints[int(s[i])]++
	}
	sort.Ints(ints)
	return ints
}
