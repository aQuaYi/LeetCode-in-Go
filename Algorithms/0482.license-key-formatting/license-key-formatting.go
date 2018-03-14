package problem0482

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	// 统计 s 中的 "-" 的个数
	countDashs := strings.Count(s, "-")
	// 统计 s 中字母和数字的个数
	countAN := len(s) - countDashs

	if countAN == 0 {
		return ""
	}

	// 删除 '-' ，并把所有的字母改成大写
	s = strings.Replace(s, "-", "", -1)
	s = strings.ToUpper(s)

	// (countAN+k-1)/k-1 是 res 中 '-' 的个数
	res := make([]byte, (countAN+k-1)/k-1+countAN)

	i, j := len(res), len(s)
	for 0 <= j-k {
		copy(res[i-k:i], s[j-k:j])

		if 0 <= i-k-1 {
			res[i-k-1] = '-'
		}

		i -= k + 1
		j -= k
	}

	// 处理头部不完整的字符
	if j > 0 {
		copy(res[:j], s[:j])
	}

	return string(res)
}
