package problem0874

var dxs = []int{0, 1, 0, -1}
var dys = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	isBlocked := make(map[int]bool, 10000)
	for _, o := range obstacles {
		i, j := o[0], o[1]
		isBlocked[encode(i, j)] = true
	}

	x, y, res := 0, 0, 0
	index := 0

	for _, c := range commands {
		switch {
		case c == -2:
			index--
		case c == -1:
			index++
		default:
			if index < 0 {
				index += 1<<63 - 4
			}
			index %= 4
			dx, dy := dxs[index], dys[index]
			for c > 0 && !isBlocked[encode(x+dx, y+dy)] {
				c--
				x += dx
				y += dy
			}
			res = max(res, x*x+y*y)
		}
	}

	return res
}

func encode(x, y int) int {
	x &= 0xFFFF
	y &= 0xFFFF
	return x<<16 | y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 以下方法太繁琐了
// 应该注意到，每次最多移动 9 步，所以，走一步看一步，也是一个好方法。

// func robotSim(commands []int, obstacles [][]int) int {

// 	sort.Slice(obstacles, func(i int, j int) bool {
// 		if obstacles[i][0] == obstacles[j][0] {
// 			return obstacles[i][1] < obstacles[j][1]
// 		}
// 		return obstacles[i][0] < obstacles[j][0]
// 	})
// 	oxs := make(map[int][]int, len(obstacles))
// 	oys := make(map[int][]int, len(obstacles))
// 	for _, o := range obstacles {
// 		if o[0] == 0 && o[1] == 0 {
// 			continue
// 		}
// 		oxs[o[0]] = append(oxs[o[0]], o[1])
// 		oys[o[1]] = append(oys[o[1]], o[0])
// 	}

// 	x, y := 0, 0

// 	north := func(a int) {
// 		y += a
// 		ys, ok := oxs[x]
// 		if ok {
// 			i := sort.SearchInts(ys, y-a)
// 			if i < len(ys) {
// 				y = min(y, ys[i]-1)
// 			}
// 		}
// 	}

// 	east := func(a int) {
// 		x += a
// 		xs, ok := oys[y]
// 		if ok {
// 			i := sort.SearchInts(xs, x-a)
// 			if i < len(xs) {
// 				x = min(x, xs[i]-1)
// 			}
// 		}
// 	}

// 	south := func(a int) {
// 		y -= a
// 		ys, ok := oxs[x]
// 		if ok {
// 			i := sort.SearchInts(ys, y+a)
// 			if 0 < i {
// 				y = max(y, ys[i-1]+1)
// 			}
// 		}
// 	}

// 	west := func(a int) {
// 		x -= a
// 		xs, ok := oys[y]
// 		if ok {
// 			i := sort.SearchInts(xs, x+a)
// 			if 0 < i {
// 				x = max(x, xs[i-1]+1)
// 			}
// 		}
// 	}

// 	moves := []func(int){north, east, south, west}
// 	idx := 40000
// 	mov := north

// 	turn := func(c int) {
// 		if c == -2 {
// 			idx--
// 		} else {
// 			idx++
// 		}
// 		mov = moves[idx%4]
// 	}

// 	// 合并同一个方向连续移动的步数
// 	tmp := make([]int, 1, len(commands))
// 	tmp[0] = commands[0]
// 	j := 0
// 	for i := 1; i < len(commands); i++ {
// 		c := commands[i]
// 		if tmp[j] > 0 && c > 0 {
// 			tmp[j] += c
// 		} else {
// 			tmp = append(tmp, c)
// 			j++
// 		}
// 	}
// 	commands = tmp

// 	res := 0

// 	for _, c := range commands {
// 		if c < 0 {
// 			turn(c)
// 		} else {
// 			mov(c)
// 			res = max(res, x*x+y*y)
// 		}
// 	}

// 	return res
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
