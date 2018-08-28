package problem0893

func numSpecialEquivGroups(A []string) int {
	strSize := len(A[0])
	groups := make(map[[26]int]bool, len(A))
	for _, a := range A {
		count := [26]int{}
		i := 0
		for i = 0; i+1 < strSize; i += 2 {
			// 使用循环展开，避免了判断 i 的奇偶性的 mod 运算
			// 还减短了关键路径的长度。
			count[a[i]-'a']++
			// 由于题目条件限制了 strSize <= 20
			// 所以可以在同一个数据同时统计奇偶位上字母
			count[a[i+1]-'a'] += 100
		}
		if i < strSize {
			count[a[i]-'a']++
		}
		groups[count] = true
	}
	return len(groups)
}
