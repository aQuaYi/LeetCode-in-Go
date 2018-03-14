package problem0744

import (
	"sort"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})
	return letters[i%n]
}
