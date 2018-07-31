package problem0866

import (
	"math"
	"strconv"
)

func primePalindrome(n int) int {
	n = nextPalindrome(n)

	for !isPrime(n) {
		n = palindromeAfterPalindrome(n)
	}

	return n
}

func isPalindrome(x int) bool {
	s := []byte(strconv.Itoa(x))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func nextPalindrome(n int) int {
	for !isPalindrome(n) {
		n++
	}
	return n
}

func palindromeAfterPalindrome(n int) int {
	if n < 9 {
		return n + 1
	}

	ok, base, length := isAll9(n)
	if ok {
		return base + 1
	}

	left := n
	rightRevers := 0
	for i := length; i > ((length / 2) + (length % 2)); i-- {
		rightRevers = rightRevers*10 + left%10
		left /= 10
	}

	if length%2 == 1 {
		rightRevers = rightRevers*10 + left%10
	}

	if left <= rightRevers {
		left++
	}

	rightRevers = left
	res := left
	if length%2 == 1 {
		rightRevers /= 10
	}
	for rightRevers > 0 {
		res = res*10 + rightRevers%10
		rightRevers /= 10
	}

	return res
}

func isAll9(n int) (bool, int, int) {
	base := 1
	len := 0
	for n%10 == 9 {
		n /= 10
		base *= 10
		len++
	}
	t := n
	for t > 0 {
		t /= 10
		len++
	}
	return n == 0, base, len
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

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}
