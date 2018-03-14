package problem0696

func countBinarySubstrings(s string) int {
	count, countZero, countOne := 0, 0, 0
	prev := rune(s[0])

	for _, r := range s {
		if prev == r {
			if r == '0' {
				countZero++
			} else {
				countOne++
			}
		} else {
			// 较少的数字决定了符合题意的子字符串个数
			// 例如 "00011" 符合题意的子字符串为 "0011", "01"，其中第一个 "0" 是无用的
			count += min(countZero, countOne)
			if r == '0' {
				countZero = 1
			} else {
				countOne = 1
			}
		}
		prev = r
	}

	return count + min(countZero, countOne)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
