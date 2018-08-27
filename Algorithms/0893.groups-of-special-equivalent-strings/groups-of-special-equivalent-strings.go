package problem0893

func numSpecialEquivGroups(A []string) int {
	aSize := len(A[0])
	gMap := make(map[[26]int]bool, len(A))
	for _, a := range A {
		count := [26]int{}
		i := 0
		for i = 0; i+1 < aSize; i += 2 {
			count[a[i]-'a']++
			count[a[i+1]-'a'] += 100
		}
		if i < aSize {
			count[a[i]-'a']++
		}
		gMap[count] = true
	}
	return len(gMap)
}
