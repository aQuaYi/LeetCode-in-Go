package problem0866

import (
	"math"
	"strconv"
)

var special = []int{2, 2, 2, 3, 5, 5, 7, 7, 11, 11, 11, 11}

func primePalindrome(N int) int {
	if N <= 11 {
		return special[N]
	}

	/** 由于 11 是偶数长度的 palindrome 的因子
	 *  所以答案一定是奇数长度的
	 *  假设答案的形式如 lmr (Left + Middle + Right) */
	lm := genLM(N)

	for {
		p := genPalindrome(lm)
		if p >= N && isPrime(p) {
			/** 需要验证 p >= N，是因为
			 * 刚开始的 p 可能会比 N 小，
			 * 例如，N = 98390 时，第一个 p 是 98389 */
			return p
		}
		lm++
	}
}

// 根据 N 生成 palindrome 的 left + middle 部分
func genLM(N int) int {
	size := len(strconv.Itoa(N))
	base := int(math.Pow10(size / 2))

	/** lm 是 N 的左边和中间部分 */
	lm := N / base

	if size&1 == 0 {
		/** 如果 N 的长度是偶数， 把 base 当做 lm
		 * 直接从 10^size+1 开始检查*/
		lm = base
	}

	return lm
}

// 利用 left 和 middle 生成一个奇数位长度的 palindrome
func genPalindrome(lm int) int {
	res := lm
	/**把 l 按照相反的顺序放到 lm 的右边，就形成了奇数位的 palindrome */
	for l := lm / 10; l > 0; l /= 10 {
		res = res*10 + l%10
	}
	return res
}

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

/**
 * 此题要点：
 * 1. 利用数学知识简化需要处理的情况
 *    形如 abccba 等长度为偶数的回文数，都不是素数，因为有因子 11
 * 2. 把 res <= 101 这样的与后面处理方式不兼容的少数情况，单独处理。
 * 3. 没有把寻找 n 作为单独的函数，简化了判断条件。
 */
