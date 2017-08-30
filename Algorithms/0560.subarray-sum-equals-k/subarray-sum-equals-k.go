package Problem0560

func subarraySum(a []int, k int) int {
	cnt := 0

	for i := 0; i < len(a); i++ {
		tmp := 0
		for j := i; j < len(a); j++ {
			tmp += a[j]
			if tmp == k {
				cnt++
			}
		}
	}

	return cnt
}
