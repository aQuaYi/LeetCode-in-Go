package problem0075

// 借鉴三路快排中的划分思路
func sortColors(a []int) {
	i, j, k := 0, 0, len(a)-1

	// for 循环中， nums[i:j] 中始终全是 1
	// 循环结束后，
	// nums[:i] 中全是 0
	// nums[j:] 中全是 2
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
