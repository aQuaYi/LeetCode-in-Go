package problem0898

func subarrayBitwiseORs(a []int) int {
	unique := make(map[int]bool, len(a))
	var prev, next []int
	for _, x := range a {
		isInNext := make(map[int]bool, len(prev))
		isInNext[x], unique[x] = true, true
		next = append(next, x)
		for _, y := range prev {
			y |= x
			if !isInNext[y] {
				isInNext[y], unique[y] = true, true
				next = append(next, y)
			}
		}
		// 假设 x 在 a 中的索引号是 j，BitwiseOR(a) 表示对 a 中所有元素依次取或的结果
		// 那么此时， next 是 BitwiseOR(a[i:j+1]) 的集合，其中 i=0,1,...,j
		prev, next = next, prev[:0]
	}
	return len(unique)
}
