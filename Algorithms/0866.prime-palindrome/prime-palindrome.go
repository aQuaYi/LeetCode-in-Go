package problem0866

import (
	"math"
	"strconv"
)

func primePalindrome(N int) int {

	switch {
	case N <= 2:
		return 2
	case N == 3:
		return 3
	case N <= 5:
		return 5
	case N <= 7:
		return 7
	case N <= 11:
		return 11
	case N <= 101:
		return 101
	}

	size := len(strconv.Itoa(N))

	left := 1
	for i := 0; i < size/2; i++ {
		left *= 10
	}
	if size%2 == 1 {
		left = N / left
	}

	// 当 size 为偶数时，
	// 可知
	// 要么，下个回文数的长度，也为偶数
	// 要么，下个回文数是 10^size + 1
	// 又
	// 当回文数的长度为偶数时，肯定不是质数，因为 11 是因子
	// 所以，直接从 10^size + 1 开始找起

	for {
		// 以 left 的末尾为镜子，制造出回文数
		// 假设 left 看起来像 abc，则要找一个 n 看起来像 abcba
		n := left
		tmp := left / 10
		for tmp > 0 {
			n = n*10 + tmp%10
			tmp /= 10
		}

		// 需要验证 n >= N，是因为
		// 刚开始的 n 可能会比 N 小，
		// 例如，N = 98390 时，第一个 n 是 98389
		if n >= N && isPrime(n) {
			return n
		}

		left++
	}
}

/**
 * 此题要点：
 * 1. 利用数学知识简化需要处理的情况
 *    形如 abccba 等长度为偶数的回文数，都不是素数，因为有因子 11
 * 2. 把 res <= 101 这样的与后面处理方式不兼容的少数情况，单独处理。
 * 3. 没有把寻找 n 作为单独的函数，简化了判断条件。
 */

// https://blog.csdn.net/willduan1/article/details/50975381
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	k := int(math.Sqrt(float64(n))) + 1
	for i := 5; i < k; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
