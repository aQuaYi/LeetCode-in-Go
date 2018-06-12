package problem0841

func canVisitAllRooms(rooms [][]int) bool {
	N := len(rooms)

	next := make([]int, 1, N)
	next[0] = 0

	isEntered := make([]bool, N)
	isEntered[0] = true
	count := 1

	for len(next) > 0 {
		r := next[0]
		next = next[1:]

		for _, x := range rooms[r] {
			if isEntered[x] {
				continue
			}
			next = append(next, x)
			isEntered[x] = true
			count++
		}

		if count == N {
			return true
		}

	}

	return count == N
}
