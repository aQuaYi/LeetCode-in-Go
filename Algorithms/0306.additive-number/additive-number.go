package Problem0306

import (
	"strconv"
)

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	var dfs func(int, int, string) bool
	dfs = func(n1, n2 int, num string) bool {
		n3 := n1 + n2
		s3 := strconv.Itoa(n3)
		if len(s3) > len(num) {
			return false
		}

		if len(s3) == len(num) {
			if s3 == num {
				return true
			}
			return false
		}

		if s3 != num[:len(s3)] {
			return false
		}

		return dfs(n2, n3, num[len(s3):])
	}

	max := len(num) * 2 / 3
	var i, j int

	for i = 1; i <= max-1; i++ {
		if len(num[:i]) > 1 && num[0] == '0' {
			return false
		}
		for j = i + 1; j <= max; j++ {
			if len(num[i:j]) > 1 && num[i] == '0' {
				break
			}
			n1, _ := strconv.Atoi(num[:i])
			n2, _ := strconv.Atoi(num[i:j])
			if dfs(n1, n2, num[j:]) {
				return true
			}
		}
	}

	return false
}
