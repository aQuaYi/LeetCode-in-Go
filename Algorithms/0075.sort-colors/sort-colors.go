package problem0075

// 借鉴三路快排中的划分思路
func sortColors(a []int) {
	size := len(a)
	if size == 0 {
		return
	}

	i, j, k := 0, 0, size-1

	// for 循环中， nums[i:j] 中始终全是 1
	for j <= k {
		switch a[j] {
		case 0:
			a[i], a[j] = a[j], a[i]
			i++
			j++
		case 1:
			j++
		case 2:
			a[j], a[k] = a[k], a[j]
			k--
		}
	}

}
