package Problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 把 string 转换成 []rune 可以适应 Unicode 字符
	sr := []rune(s)
	tr := []rune(t)

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
