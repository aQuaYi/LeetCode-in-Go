package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 把 string 转换成 []rune 可以适应 Unicode 字符
	sr := []rune(s)
	tr := []rune(t)

	// 因为使用了 []rune，rec 只好使用 map
	// 不然的话，使用 [26]int 数组，效率更高
	rec := make(map[rune]int, len(sr))
	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	for _, n := range rec {
		if n != 0 {
			return false
		}
	}

	return true
}
