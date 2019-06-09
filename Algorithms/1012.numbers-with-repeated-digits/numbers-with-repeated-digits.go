package problem1012

// ref: https://leetcode.com/problems/numbers-with-repeated-digits/discuss/256725/JavaPython-Count-the-Number-Without-Repeated-Digit
func numDupDigitsAtMostN(N int) int {
	digits := convert(N + 1) // +1 为为了把判别条件从 <= 简化成 <
	bits := len(digits)

	noRepeat := 0
	// 统计宽度短于 N 的不含重复数字的数的个数
	for b := 1; b < bits; b++ {
		noRepeat += 9 * count(9, b-1)
	}

	// 统计和 N 具有相同高位的且不含重复数字的数的个数
	hasSeen := make(map[int]bool, 10)
	for b := 0; b < bits; b++ {
		d := 0
		if b == 0 {
			d = 1 // 最高位上的数字不能为 0
		}
		for ; d < digits[b]; d++ {
			// 如果 hasSeen[d]=true
			// 说明在更高位上已经出现过 d 了
			// 无论低位上再出现什么，都不是 noRepeat 了
			if !hasSeen[d] {
				noRepeat += count(9-b, bits-b-1)
			}
		}
		if hasSeen[digits[b]] {
			// digits[:b+1] 中出现了重复的数字
			// 不可能再有 noRepeat 了
			break
		}
		hasSeen[digits[b]] = true
	}
	return N - noRepeat
}

func convert(N int) []int {
	res := make([]int, 0, 10)
	for N > 0 {
		res = append(res, N%10)
		N /= 10
	}
	swap(res)
	return res
}

func swap(A []int) {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}

// 统计了 m 个数字组成宽度为 n 的数值时，
// 不含重复数字的数值的个数。
// 注意，
// 真正使用 count 的时候，是为了统计 n+1 位宽度的数值
// 在 n+1 位的最高位上不放置 0
// 就可以保证剩下的 n 位的宽度确实是 n
// 所以，count 在使用时，还需要乘上 n+1 位上可供选择的数字的个数
// 例如， 9*count(9, i-1)  的实际含义才是，所有宽度为 i 且不包含重复数字的数值个数
// 例如， count（8，i） 的实际含义是，以 X 开头所有宽度为 i+1 且不包含重复数字的数的个数
//      X 可以是任意一个数字
func count(m, n int) int {
	if n == 0 {
		return 1
	}
	return count(m, n-1) * (m - n + 1)
}
