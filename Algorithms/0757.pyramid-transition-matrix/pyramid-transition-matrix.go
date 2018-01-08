package Problem0757

import (
	"strings"
)

func pyramidTransition(bottom string, allowed []string) bool {
	blocks := "#" + strings.Join(allowed, "#")

	return helper(bottom, blocks, 0, len(bottom))
}

func helper(bottom, blocks string, curr, length int) bool {
	if curr+2 > length {
		bottom = bottom[length:]
		length = len(bottom)
		if length == 1 {
			return true
		}
		return helper(bottom, blocks, 0, length)
	}

	b := "#" + bottom[curr:curr+2]

	count := strings.Count(blocks, b)

	for i := 0; i < count; i++ {
		index := strings.Index(blocks, b)
		blk := blocks[index : index+4]
		color := blk[3:]
		blocks = strings.Replace(blocks, blk, "", 1)
		blocks += blk
		if helper(bottom+color, blocks, curr+1, length) {
			return true
		}
	}

	return false
}
