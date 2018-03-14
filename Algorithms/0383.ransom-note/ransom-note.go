package problem0383

func canConstruct(ransomNote string, magazine string) bool {
	mc := getCount(magazine)

	for _, b := range ransomNote {
		mc[b-'a']--
		if mc[b-'a'] < 0 {
			return false
		}
	}

	return true
}

func getCount(s string) []int {
	res := make([]int, 26)
	for i := range s {
		res[s[i]-'a']++
	}
	return res
}
