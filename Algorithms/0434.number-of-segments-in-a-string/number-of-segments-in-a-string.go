package problem0434

import (
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	strs := strings.Fields(s)
	return len(strs)
}
