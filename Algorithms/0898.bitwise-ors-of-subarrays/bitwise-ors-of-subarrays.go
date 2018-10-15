package problem0898

func subarrayBitwiseORs(a []int) int {
	unique := make(map[int]bool, len(a))
	var s0, s1 []int
	for _, x := range a {
		tmp := make(map[int]bool, len(s0))
		tmp[x], unique[x] = true, true
		s1 = append(s1, x)
		for _, y := range s0 {
			y |= x
			if !tmp[y] {
				tmp[y], unique[y] = true, true
				s1 = append(s1, y)
			}
		}
		s0, s1 = s1, s0[:0]
	}
	return len(unique)
}
