package Problem0400

import "math"

func findNthDigit(n int) int {
	var i, pre, prelen int
	for t := n; t > 0; i++ {
		prelen = findRangeMaxLen(i + 1)
		pre += prelen
		t -= prelen
	}
	pre -= prelen

	extra := n - pre

	// i 代表解在第几位里

	extraH := extra / i
	left := extra - (extraH * i)
	// println("extraH", extraH, left)

	last := int(math.Pow10(i-1)) + extraH - 1
	if left != 0 {
		last++
	} else {
		left = i
	}

	return nthChar(last, left, i)
}

func findRangeMaxLen(r int) int {
	return 9 * int(math.Pow10(r-1)) * r
}

func nthChar(num, index, len int) int {
	for i := (len - index); i > 0; i-- {
		num /= 10
	}

	return num - num/10*10
}
