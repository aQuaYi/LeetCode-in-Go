package problem0409

func longestPalindrome(s string) int {
	res := 0
	a := [123]int{} // 'z' 的 ASCII 码为 122
	for i := range s {
		a[s[i]]++
	}

	// hasOdd 表示存在数目为奇数的元素，可以放在中间
	hasOdd := 0
	for i := range a {
		if a[i] == 0 {
			continue
		}
		if a[i]&1 == 0 {
			res += a[i]
		} else {
			res += a[i] - 1
			hasOdd = 1
		}
	}

	return res + hasOdd
}
