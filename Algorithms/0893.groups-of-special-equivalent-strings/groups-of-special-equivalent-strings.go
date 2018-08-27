package problem0893

func numSpecialEquivGroups(A []string) int {
	aSize := len(A[0])
	oddBit := 1 << 31
	gMap := make(map[[26]int]bool, len(A))
	for _, a := range A {
		count := [26]int{}
		i := 0
		for i = 0; i+1 < aSize; i += 2 {
			// 使用循环展开，避免了判断 i 的奇偶性的 mod 运算
			// 还减短了关键路径的长度。
			count[a[i]-'a']++
			// 由于题目条件限制了 aSize <= 20
			// 所以可以利用同一个数据同时统计奇偶位上字母
			count[a[i+1]-'a'] += oddBit
		}
		if i < aSize {
			count[a[i]-'a']++
		}
		gMap[count] = true
	}
	return len(gMap)
}
