package problem0893

func numSpecialEquivGroups(A []string) int {
	gMap := make(map[int]int, len(A))
	for _, a := range A {
		gMap[encode(a)]++
	}
	return len(gMap)
}

func encode(str string) int {
	evens, odds := 0, 0
	for i, r := range str {
		if i%2 == 0 {
			evens += int(r)
		} else {
			odds += int(r)
		}
	}
	return evens<<16 | odds
}
