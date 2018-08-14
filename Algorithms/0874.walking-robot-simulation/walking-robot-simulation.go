package problem0874

import "sort"

func robotSim(commands []int, obstacles [][]int) int {

	sort.Slice(obstacles, func(i int, j int) bool {
		if obstacles[i][0] == obstacles[j][0] {
			return obstacles[i][1] < obstacles[j][1]
		}
		return obstacles[i][0] < obstacles[j][0]
	})
	oxs := make(map[int][]int, len(obstacles))
	oys := make(map[int][]int, len(obstacles))
	for _, o := range obstacles {
		if o[0] == 0 && o[1] == 0 {
			continue
		}
		oxs[o[0]] = append(oxs[o[0]], o[1])
		oys[o[1]] = append(oys[o[1]], o[0])
	}

	x, y := 0, 0

	north := func(a int) {
		ys, ok := oxs[x]
		i := sort.SearchInts(ys, y)
		y += a
		if ok && i < len(ys) {
			y = min(y, ys[i]-1)
		}
	}

	east := func(a int) {
		xs, ok := oys[y]
		i := sort.SearchInts(xs, x)
		x += a
		if ok && i < len(xs) {
			x = min(x, xs[i]-1)
		}
	}

	south := func(a int) {
		ys, ok := oxs[x]
		i := sort.SearchInts(ys, y)
		y -= a
		if ok && 0 < i {
			y = max(y, ys[i-1]+1)
		}
	}

	west := func(a int) {
		xs, ok := oys[y]
		i := sort.SearchInts(xs, x)
		x -= a
		if ok && 0 < i {
			x = max(x, xs[i-1]+1)
		}
	}

	moves := []func(int){north, east, south, west}
	idx := 40000
	mov := north

	turn := func(c int) {
		if c == -2 {
			idx--
		} else {
			idx++
		}
		mov = moves[idx%4]
	}

	tmp := make([]int, 1, len(commands))
	tmp[0] = commands[0]
	j := 0
	for i := 1; i < len(commands); i++ {
		c := commands[i]
		if tmp[j] > 0 && c > 0 {
			tmp[j] += c
		} else {
			tmp = append(tmp, c)
			j++
		}
	}

	commands = tmp

	res := 0

	for _, c := range commands {
		if c < 0 {
			turn(c)
		} else {
			mov(c)
			res = max(res, x*x+y*y)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
