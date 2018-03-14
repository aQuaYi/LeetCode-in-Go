package problem0306

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

	// n1 = num[:i]
	// n2 = num[i:j]
	// n3 = num[j:j+x], x = 1,2,3..
	var i, j int
	// Jmax 是 j 的最大值
	// n1, n2 中较大数的 size 的最小值为 (j+1)/2
	// n3 的 size 的最大值为 len(num)-j
	// 当 (j+1)/2 <= len(num)-j 才有可能使得 n1+n2==n3 成立
	// 可得 j <= (2* len(num)-1)/3 < len(num) * 2 / 3
	// 考虑到 int 除法的特性
	// Jmax = len(num) * 2 / 3
	Jmax := len(num) * 2 / 3

	for i = 1; i <= Jmax-1; i++ {
		if len(num[:i]) > 1 && num[0] == '0' {
			return false
		}
		for j = i + 1; j <= Jmax; j++ {
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
