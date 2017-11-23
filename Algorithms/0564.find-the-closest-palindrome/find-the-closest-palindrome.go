package Problem0564

import "strconv"

func nearestPalindromic(n string) string {
	if !isPalindrome(n) {
		return palindromefy(n)
	}

	diff := 1
	if len(n)%2 == 0 {
		diff = 11
	}
	i := len(n)/2 + 1
	for i < len(n) {
		diff *= 10
		i++
	}

	num, _ := strconv.Atoi(n)

	nb := num + diff
	ns := num - diff

	nbpStr := palindrome2(strconv.Itoa(nb), min)
	nbp, _ := strconv.Atoi(nbpStr)

	nspStr := palindrome2(strconv.Itoa(ns), max)
	nsp, _ := strconv.Atoi(nspStr)

	avg := (nbp + nsp) / 2

	if avg < num {
		return nbpStr
	}
	return nspStr
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func palindromefy(n string) string {
	bytes := []byte(n)
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {
			bytes[j] = bytes[i]
		}
		i++
		j--
	}
	return string(bytes)
}

func palindrome2(n string, f func(byte, byte) byte) string {
	bytes := []byte(n)
	i, j := 0, len(bytes)-1
	for i < j {
		if bytes[i] != bytes[j] {
			bytes[i] = f(bytes[i], bytes[j])
			bytes[j] = f(bytes[i], bytes[j])
		}
		i++
		j--
	}
	return string(bytes)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}
