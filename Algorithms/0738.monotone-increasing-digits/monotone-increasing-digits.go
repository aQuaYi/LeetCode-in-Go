package problem0738

import (
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	for i := len(s) - 2; 0 <= i; i-- {
		if s[i] <= s[i+1] {
			continue
		}

		// 当 s[i] > s[i+1] 时
		s[i]--
		// s[i+1:] 中的数字全部变成　９
		for j := i + 1; j < len(s); j++ {
			s[j] = '9'
		}
	}

	res, _ := strconv.Atoi(string(s))
	return res
}
