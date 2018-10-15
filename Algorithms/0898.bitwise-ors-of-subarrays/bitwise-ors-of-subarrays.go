package problem0898

func subarrayBitwiseORs(A []int) int {
	size := len(A)
	res := make(map[int]bool, size)
	var cur map[int]bool

	for _, n := range A {
		cur2 := make(map[int]bool, 30)
		cur2[n] = true
		for k := range cur {
			cur2[n|k] = true
		}
		cur = cur2
		for k := range cur {
			res[k] = true
		}
	}
	return len(res)
}
