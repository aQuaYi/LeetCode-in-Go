package Problem0553

import (
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	return strs[0] + "/(" + strings.Join(strs[1:], "/") + ")"
}
