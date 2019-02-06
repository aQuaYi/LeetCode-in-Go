package problem0943

import (
	"fmt"
	"strings"
)

func shortestSuperstring(A []string) string {
	size := len(A)
	res := strings.Repeat("?", 12*20+1)
	isUsed := make([]bool, size)
	rescur(A, isUsed, size, "", &res)
	fmt.Println(res)
	return res
}

func rescur(A []string, isUsed []bool, countDown int, tmp string, res *string) {
	if countDown == 0 {
		if len(*res) > len(tmp) {
			*res = tmp
		}
		return
	}

	for i, str := range A {
		if isUsed[i] {
			continue
		}
		isUsed[i] = true

		j := len(str)
		for !strings.HasSuffix(tmp, str[:j]) {
			j--
		}

		rescur(A, isUsed, countDown-1, tmp+str[j:], res)

		isUsed[i] = false
	}

	return
}
