package Problem0757

import (
	"strings"
)

func pyramidTransition(bottom string, allowed []string) bool {
	blocks := "#" + strings.Join(allowed, "#")

	var helper func(string, string, int, int) bool
	helper = func(bottom, blocks string, curr, length int) bool {
		if curr+2 > length {
			bottom = bottom[length:]
			length = len(bottom)
			if length == 1 {
				return true
			}
			return helper(bottom, blocks, 0, length)
		}

		b := "#" + bottom[curr:curr+2]
		beg := 0
		for beg < len(blocks) {
			index := strings.Index(blocks[beg:], b) + beg
			if index < beg {
				break
			}
			beg = index + 4
			color := blocks[index+3 : index+4]
			if helper(bottom+color, blocks, curr+1, length) {
				return true
			}
		}

		return false
	}

	return helper(bottom, blocks, 0, len(bottom))
}
