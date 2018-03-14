package problem0275

// a 为升序排列
func hIndex(a []int) int {
	size := len(a)

	// 二分查找法
	lo, hi := 0, size-1
	// lo, miD, hi 都是降序切片 d 中的序列号
	// 因为 a 是 d 的逆序，即 a 是升序切片
	// d[miD] , a[miA] 是同一个数
	// 所以，存在数量关系，miD + miA +1 == size
	var miD, miA int
	for lo <= hi {
		miD = (lo + hi) / 2
		miA = size - miD - 1
		if a[miA] > miD {
			lo = miD + 1
		} else {
			hi = miD - 1
		}
	}

	return lo
}
